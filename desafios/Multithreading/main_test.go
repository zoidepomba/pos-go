package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestViaCEP(t *testing.T) {
	mockResponse := StructViaCEP{
		Cep:         "38412-392",
		Logradouro:  "Avenida Jerusalém",
		Complemento: "",
		Bairro:      "Jardim Canaã",
		Localidade:  "Uberlândia",
		Uf:          "MG",
		Ibge:        "3170206",
		Gia:         "",
		Ddd:         "34",
		Siafi:       "5403",
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	start := time.Now()
	defer mockServer.Close()

	client := mockServer.Client()
	ch := make(chan *StructViaCEP)

	go ViaCEP(client, "38412392", ch, start)

	select {
	case result := <-ch:
		// fmt.Println(reflect.TypeOf(*result))
		// fmt.Println(reflect.TypeOf(mockResponse))
		fmt.Println(reflect.DeepEqual(result, mockResponse))
		if result == nil {
			t.Fatalf("Expected non-nil result")
		}
		if !reflect.DeepEqual(*result, mockResponse) {
			// fmt.Println(mockResponse)
			// fmt.Println(*result)
			t.Fatal("content not expected")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Test timed out")
	}
}
