package main

import (
	"net/http"
)

type RegisterPage struct{
        Name            string
        Title           string
        MenuActive      string
        HostName        string
        Errors  map[string]string
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("The is not a love song"))
	title := TITLE+" - Register"
        p := RegisterPage{Title:title, HostName:r.Host, MenuActive:"user"}
        render(w, "register", p)

}
