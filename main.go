package main

import (
	"encoding/json"
	"log"
	"net/http"
	gcp "scheduler/gcp"
	utils "scheduler/utils"
)

const port = ":5432"

func execCmd(w http.ResponseWriter, r *http.Request) {
	instance := r.PathValue("inst_name")
	values := gcp.Action(r.PathValue("action"))

	if values == nil {
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	gcp.CallGCP(values, instance)

	res_ := utils.S200("Action sent to gcp, may take a few minutes to apply")
	a, err := json.Marshal(res_)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(a)

}

func health(w http.ResponseWriter, r *http.Request) {
	pa := utils.S200("up")
	log.Println(pa)
	j, err := json.Marshal(pa)
	if err != nil {
		log.Println("Something is wrong ...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	log.Println("Everything is fine ...")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{action}/{inst_name}", execCmd)
	mux.HandleFunc("GET /health", health)

	log.Println("Listening ...")
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
