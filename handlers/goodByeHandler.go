package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	thisLogger *log.Logger
}

func NewGoodBye(thisLogger *log.Logger) *GoodBye {
	return &GoodBye{thisLogger}
}

func (h *GoodBye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Welcome to Yet another GoodBye Handler.")

	responseWriter.Write([]byte("GoodBye for today's Blog Guys."))

}
