package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type USDResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

var db *gorm.DB

func StartServer() {
	db = InitDB()

	http.HandleFunc("/cotacao", handler)
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Contexto API externa (200ms)
	ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(
		ctxAPI,
		"GET",
		"https://economia.awesomeapi.com.br/json/last/USD-BRL",
		nil,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Timeout API:", err)
		http.Error(w, "Erro ao buscar cotação", http.StatusGatewayTimeout)
		return
	}
	defer resp.Body.Close()

	var data USDResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Contexto banco (10ms)
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	if err := db.WithContext(ctxDB).Create(&Cotacao{
		Bid: data.USDBRL.Bid,
	}).Error; err != nil {
		log.Println("Timeout DB:", err)
	}

	// Retorno exigido: JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": data.USDBRL.Bid,
	})
}
