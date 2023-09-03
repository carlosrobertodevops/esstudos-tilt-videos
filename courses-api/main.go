package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	Name string `json:"name"`
}

func main(){
	courses := []Course{
		{Name: "Pós-graduação em Golang"},
		{Name: "Full Experience Cycle"},
		{Name: "Full Cycle"},
		{Name: "MBA em Arquitetura full Cycle"},
	}

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Controle-Allow-Origin", "*")
		w.Header().Set("Access-Controle-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		json.NewEncoder(w).Encode(courses)
	})
	http.ListenAndServe(":8080", nil)
}