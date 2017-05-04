package main

import(
    "net/http"
    "log"
    "fmt"
    "../mux"
)

func hola (w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w, "Hola desde una funcion anonima")
}

type User struct{
    name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola usuario " + this.name)
}

func main() {
    
    user := &User{"Dani"}
    
    mux := mux.CreateMux()
    mux.AddFunc("/hola", hola)
    mux.AddHandle("/user", user)
    
    log.Fatal(http.ListenAndServe("localhost:8000",mux))

}