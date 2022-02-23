package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	thisLogger *log.Logger
}

func NewHello(thisLogger *log.Logger) *Hello {
	return &Hello{thisLogger}
}

func (h *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Hello World")

	dataBody, error := ioutil.ReadAll(request.Body)

	if error != nil {
		/*responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Some error have come in."))
		return*/
		http.Error(responseWriter, "Some error have come in.", http.StatusBadRequest)
		return
	}

	//log.Printf("Data %s\n", dataBody)

	fmt.Fprintf(responseWriter, "Data thus received in Request is : %s\n", dataBody)
}
