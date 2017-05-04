//web server

package main

import(
    //"fmt"
    "log"
    "net/http"
)

func main(){
    http.HandleFunc("/", handler)//cada peticion manda llamar un handler
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//handler imprime el componente Path de la peticion URL r
func handler(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    
    //dos maneras de hacer redirect
    //http.Redirect(w, r, "https://www.google.com", 301)
    http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently)
}