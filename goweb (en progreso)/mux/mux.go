package mux

import(
    "net/http"
    
)

type myHandler func(http.ResponseWriter, *http.Request)

type MuxDany struct{
    danyRutas map[string] myHandler//handlers 
}

func (this *MuxDany) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if fn, ok := this.danyRutas[r.URL.Path]; ok{
        fn(w,r)
    }else{
        http.NotFound(w,r)
    }
    
}

func (this *MuxDany) AddFunc(ruta string, fn myHandler){
    this.danyRutas[ruta] = fn
}

func (this *MuxDany) AddHandle(ruta string, handle http.Handler){
    this.danyRutas[ruta] = handle.ServeHTTP
}

func CreateMux() *MuxDany{
    newMap := make(map[string]myHandler)
    return &MuxDany{newMap}
}