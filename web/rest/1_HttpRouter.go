package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET(`/`, Index)
	router.GET(`/hello/:name`, Hello)
	router.GET(`/getuser/:uid`, getuser)
	router.GET(`/deleteuser/:uid`, deleteuser)

	log.Fatal(http.ListenAndServe(":9999", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := p.ByName("uid")
	fmt.Fprintf(w, "you are added user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are deleted user %s", uid)
}
