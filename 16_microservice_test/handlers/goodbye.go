package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoodByeStruct does something
type GoodByeStruct struct {
	l *log.Logger
}

// NewGoodBye does something
func NewGoodBye(l *log.Logger) *GoodByeStruct {
	return &GoodByeStruct{l}
}

func (g *GoodByeStruct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
	s := fmt.Sprintf("From goodbye, data passed in via request: '%s'", string(bs))
	g.l.Println(s)
	rw.Write([]byte(s))
}
