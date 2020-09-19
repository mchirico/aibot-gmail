package http

import (
	"github.com/mchirico/aibot-gmail/pkg/http/handles"
	"log"
	"net/http"
	_ "time/tzdata"
)

func SetupHandles() {

	http.HandleFunc("/status", handles.Status)
	http.HandleFunc("/", handles.BaseRoot)
	http.HandleFunc("/subscript", handles.Subscript)

	//http.Handle("/static/", http.StripPrefix("/static", fs))

}

func Server() {
	//go gmail.RunEmail()
	SetupHandles()
	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
