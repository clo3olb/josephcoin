package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/clo3olb/josephcoin/blockchain"
)

var templates *template.Template

const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/add", add)
	handler.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
