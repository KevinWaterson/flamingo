package main

import (
	"net/http"
)

type IndexPage struct{
        Name            string
        Title           string
        MenuActive      string
        HostName        string
        Errors  map[string]string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("The is not a love song"))
        p := IndexPage{Title:TITLE, HostName:r.Host, MenuActive:"index"}
        render(w, "index", p)

}
