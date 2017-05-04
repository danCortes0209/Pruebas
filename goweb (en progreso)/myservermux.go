package main

import(
    "net/http"
    "log"
    "fmt"
)

type myHandler func(http.ResponseWriter, *http.Request)

type MuxDany struct{
    danyRutas map[string] myHandler//handlers 
}

func (this *MuxDany) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fn := this.danyRutas[r.URL.Path]
    fn(w,r)
}

func (this *MuxDany) AddMux(ruta string, fn myHandler){
    this.danyRutas[ruta] = fn
}

func main() {
    
    newMap := make(map[string]myHandler)
    mux := &MuxDany{danyRutas: newMap}
    mux.AddMux("/hola", func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintf(w, "Hola desde una funcion anonima")
    })
    
    log.Fatal(http.ListenAndServe("localhost:8000",mux))
/*
    redirect := http.RedirectHandler("https://google.com.mx",http.StatusMovedPermanently)
    notFound := http.NotFoundHandler()
    
    http.Handle("/redirect",redirect)
    http.Handle("/not",notFound)
    
    server := &http.Server{
        Addr: "localhost:8000",
        Handler: nil,
    }
    
    log.Fatal(server.ListenAndServe())
*/
}
