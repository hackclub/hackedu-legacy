package hackedu

import (
	"fmt"
	"net/http"

	"github.com/realistschuckle/gohaml"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	engine, err := gohaml.NewEngine("%h1 Hello, world!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, engine.Render(nil))
}
