/*
	CLeyton Henrique

#API Restful em Go
	*produto
	*codigo
    *nome
    *valor

#Linha de código para inserção no linux
	*go get github.com/gorilla/mux

	O pacote gorilla/mux implementa um roteador e distribuidor de
	solicitação para atender solicitações de entrada ao respectivo manipulador.

#Linha de código para testar o código

	$go build API_Restiful.go
	$./API_Restiful

*/

package main

import (
	"encoding/json" //codificação e decodificação de json

	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type Produto struct {
	ID     string  `json:"id,omitempty"`
	Codigo int64   `json:"codigo,omitempty"`
	Nome   string  `json:"nome,omitempty"`
	Valor  float64 `json:"valor,omitempty"`
}

var produto []Produto

// Ponto final
func GetProdEndpoint(w http.ResponseWriter, req *http.Request) {
	parametro := mux.Vars(req)
	for _, item := range produto {
		if item.ID == parametro["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Produto{})
}

func GetProdutoEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(produto)
}

func CreateProdEndpoint(w http.ResponseWriter, req *http.Request) {
	parametro := mux.Vars(req)
	var prod Produto
	_ = json.NewDecoder(req.Body).Decode(&prod)
	prod.ID = parametro["id"]
	produto = append(produto, prod)
	json.NewEncoder(w).Encode(produto)

}

func DeleteProdEndpoint(w http.ResponseWriter, req *http.Request) {
	parametro := mux.Vars(req)
	for index, item := range produto {
		if item.ID == parametro["id"] {
			produto = append(produto[:index], produto[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(produto)
}

func main() {

	router := mux.NewRouter()

	produto = append(produto, Produto{ID: "1", Codigo: 767434, Nome: "Hd Externo", Valor: 300.00})
	produto = append(produto, Produto{ID: "2", Codigo: 647326, Nome: "Mouse Wireless", Valor: 59.00})
	produto = append(produto, Produto{ID: "3", Codigo: 872347, Nome: "Kit Teclado e Mouse Wireless", Valor: 144.00})
	produto = append(produto, Produto{ID: "4", Codigo: 732644, Nome: "Fone de Ouvido Bluetooth", Valor: 199.00})

	// endpoints
	router.HandleFunc("/produto", GetProdutoEndpoint).Methods("GET")
	router.HandleFunc("/produto/{id}", GetProdEndpoint).Methods("GET")
	router.HandleFunc("/produto/{id}", CreateProdEndpoint).Methods("POST")
	router.HandleFunc("/produto/{id}", DeleteProdEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
