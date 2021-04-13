package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HellorStruct does something
type HellorStruct struct {
	l *log.Logger
}

// NewHello does something
func NewHello(l *log.Logger) *HellorStruct {
	return &HellorStruct{l}
}

func (h *HellorStruct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "From hello, data passed in via request: '%s'", string(d))
	rw.WriteHeader(http.StatusAccepted)
}
