package main

import (
	"encoding/json"
	"exec/go/exec/blockchain/core"
	"net/http"
)

type BlockChainResponse struct {
	BlockChain *core.BlockChain
	Total      int
}

var bc *BlockChainResponse

func main() {
	bc = &BlockChainResponse{BlockChain: core.NewBlockChain()}
	bc.Total = len(bc.BlockChain.Blocks)
	run()
}

func run() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.ListenAndServe(":8888", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	bc.BlockChain.SendData(data)
	bc.Total = len(bc.BlockChain.Blocks)
	get(w, r)
}

func get(w http.ResponseWriter, r *http.Request) {
	//bytes, _ := json.Marshal(bc)

	_ = json.NewEncoder(w).Encode(bc)
}
