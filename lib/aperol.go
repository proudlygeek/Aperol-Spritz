package aperol

import (
    "net/http"
    "fmt"
)

type Handler struct {}

func (this *Handler) Get(route string, handle func() string) {
    http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, handle())
    })
}

func (this *Handler) Run(port string) {
    fmt.Println("Running on port " + port)
    http.ListenAndServe(port, nil)
}

func Spritz() *Handler {
    return &Handler{}
}
