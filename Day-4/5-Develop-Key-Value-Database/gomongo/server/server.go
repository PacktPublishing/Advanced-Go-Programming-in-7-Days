package main

import (
	"os"
	"net/http"
	"os/signal"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-4/5-Develop-Key-Value-Database/gomongo"
	"strings"
	"io/ioutil"
	"encoding/json"
)

var db *gomongo.Driver

func init()  {
	db, _ = gomongo.New("db")
}

const listenPort = ":3010"

func main() {

	http.HandleFunc("/gomongo/", requestHandler)
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		db.Close()
		os.Exit(0)
	}()

	http.ListenAndServe(listenPort, nil)
}

func requestHandler(rw http.ResponseWriter, req *http.Request) {
	urlPart := strings.Split(req.URL.Path, "/")
	var err error
	var document string
	var resource string

	if len(urlPart) == 4 {
		document = urlPart[2]
		resource = urlPart[3]
	}
	if len(urlPart) == 3 {
		document = urlPart[2]
	}

	switch req.Method {
	case http.MethodGet:
		var v *json.RawMessage
		err := db.Read(document, resource, &v)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.Write(*v)
		}
	case http.MethodDelete:
		err = db.Delete(document, resource)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	case http.MethodPut:
		v, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}

		err = db.Write(document, resource, v)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	default:
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
}
