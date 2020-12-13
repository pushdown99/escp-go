package main

import (
	"github.com/gorilla/mux"
    "html/template"
	"net/http"
	"log"
	//"fmt"
	"encoding/json"
	"io/ioutil"
)

//////////////////////////////////////////////////////////////

type Setting struct {
	Name  		string 
	Register	string
	Address		string
	POS			string
	Printer		string
    Using 		bool
}

var data = Setting {
	Name: 		"Input name",
	Register: 	"Input business register number",
	Address: 	"Input address",
	POS: 		"COM1",
	Printer: 	"COM2",
	Using: 		true,
};

//////////////////////////////////////////////////////////////

func Index (w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func Get (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func Upload (w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//name := params["name"]
	data.Name 		= r.FormValue("name")
	data.Register 	= r.FormValue("register")
	data.Address 	= r.FormValue("address")
	data.POS 		= r.FormValue("pos")
	data.Printer 	= r.FormValue("printer")
	data.Using      = r.FormValue("using") == "on"
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
	http.Redirect(w, r, "http://"+r.Host+"/", http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter ()
    r.HandleFunc("/", Index)
    r.HandleFunc("/upload", Upload).Methods("POST")
	//rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Methods("GET")

	http.Handle("/", r)
	port := ":80"
	log.Println("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, nil)) 
}