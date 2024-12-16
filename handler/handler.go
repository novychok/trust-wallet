package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/novychok/trustwallet/service"
)

type handler struct {
	parserService service.Parser
}

func (h *handler) GetBlock(w http.ResponseWriter, r *http.Request) {

	parsedBlock, err := h.parserService.GetCurrentBlock()
	if err != nil {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "error to get the current block"})
		return
	}

	wg := &sync.WaitGroup{}

	for _, tx := range parsedBlock.Result.Transactions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := h.parserService.Subscribe(tx.From, tx.To, tx)
			if err != nil {
				writeJSON(w, http.StatusBadRequest,
					map[string]string{"error": "error while subscribe the address"})
				return
			}
		}()
	}
	wg.Wait()

	addrs := h.parserService.GetAddresses()

	writeJSON(w, http.StatusBadRequest, addrs)
}

func (h *handler) GetTransactions(w http.ResponseWriter, r *http.Request) {

	addr := r.URL.Query().Get("address")
	if addr == "" {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "error to get the query param"})
		return
	}

	transactions, err := h.parserService.GetTransactions(addr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "error to get the transactions"})
		return
	}

	writeJSON(w, http.StatusOK, transactions)
}

func writeJSON(w http.ResponseWriter, status int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, err := json.MarshalIndent(msg, " ", " ")
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func New(parserService service.Parser) *handler {
	return &handler{
		parserService: parserService,
	}
}
