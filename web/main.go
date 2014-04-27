package main

import (
	"log"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles(
        "./"+TEMPLATE_DIR+"/header.html",
        "./"+TEMPLATE_DIR+"/footer.html",
        "./"+TEMPLATE_DIR+"/login.html",
        "./"+TEMPLATE_DIR+"/admin.html",
        "./"+TEMPLATE_DIR+"/admin_header.html",
        "./"+TEMPLATE_DIR+"/admin_footer.html",
        "./"+TEMPLATE_DIR+"/register.html",
        "./"+TEMPLATE_DIR+"/index.html"))


func logHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v: %v", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func main() {
	// serve files from the assets directory (css/js/less/img)
	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	http.Handle("/", logHandler( http.HandlerFunc( indexHandler ) ) )

	// user handlers
	http.Handle("/user/login", logHandler( http.HandlerFunc( loginHandler ) ) )
	http.Handle("/user/register", logHandler( http.HandlerFunc( registerHandler ) ) )
	// http.Handle("/group/add", logHandler( http.HandlerFunc( groupAddHandler ) ) )


	// admin handlers
	http.Handle("/admin", logHandler( http.HandlerFunc( adminHandler ) ) )


	log.Println("Listening on port: "+HTTP_PORT)
	http.ListenAndServe(":"+HTTP_PORT, nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        err := templates.ExecuteTemplate(w, filename, data)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

