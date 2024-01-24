package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strings"
)

type dump struct {
	Host   string              `json:"host"`
	Header map[string][]string `json:"headers"`
	Path   string              `json:"path"`
	Query  map[string]string   `json:"query"`
	Env    map[string]string   `json:"env"`
}

var env bool

func init() {
	flag.BoolVar(&env, "e", false, "include os env in dump")
	flag.Parse()
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump := &dump{
			Host:   r.Host,
			Header: r.Header,
			Path:   r.URL.Path,
			Query:  map[string]string{},
		}
		for k, v := range r.URL.Query() {
			dump.Query[k] = v[0]
		}
		if env {
			dump.Env = map[string]string{}
			for _, v := range os.Environ() {
				k := v[:len(v)-len(v[strings.Index(v, "="):])]
				dump.Env[k] = os.Getenv(k)
			}
		}
		data, err := json.Marshal(dump)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
