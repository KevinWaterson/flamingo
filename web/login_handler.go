package main

import (
	"net/http"
)

type LoginPage struct{
        Name            string
        Title           string
        MenuActive      string
        HostName        string
        Errors  map[string]string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("The is not a love song"))
	title := TITLE+" - login"
        p := LoginPage{Title:title, HostName:r.Host, MenuActive:"user"}
        render(w, "login", p)

}
