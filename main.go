package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/clo3olb/josephcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	data := homeData{"home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
	fmt.Fprint(rw, "Hello from home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
