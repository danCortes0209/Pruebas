package main

import (
    "net/http"
    //"net/url"
    "fmt"
    //"io/ioutil"
    "log"
)

/*
func createURL() string{
    u,err := url.Parse("/params")
    if err!=nil{
        panic(err)
    }
    
    u.Host = "localhost:8000"
    u.Scheme = "http"
    
    query := u.Query()
    query.Add("nombre","valor")
    u.RawQuery = query.Encode()
    
    return u.String()
}

func hola(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola Mundo")
}

func holados(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola Mundo v2")
}
*/

type User struct{
    name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola: " + this.name)
}

func main(){
    
    //mux := http.NewServeMux()
    //mux.HandleFunc("/", hola)
    //mux.HandleFunc("/dos",holados)
    
    dan := &User{name: "Dano"}
    
    mux := http.NewServeMux()
    mux.Handle("/dan", dan)
    
    server := &http.Server{
        Addr: "localhost:8000",
        Handler: mux,//si es nil se usa el Default ServerMux
    }
    
    /*
    url := createURL()
    request, err := http.NewRequest("GET", url, nil)
    if err!=nil{
        panic(err)
    }
    
    client := &http.Client{}
    req,eror := client.Do(request)
    
    if eror!=nil{
        panic(eror)
    }
    
    fmt.Println("Header", req.Header)
    body,err := ioutil.ReadAll(req.Body)
    if err!= nil{
        panic(err)
    }
    fmt.Println("body",string(body))
    fmt.Println("Status",req.Status)
    
    fmt.Println("la url final es: " + url)
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Println(r.Header)
        
        accessToken := r.Header.Get("access_token")
        if len(accessToken) != 0{
            fmt.Fprintf(w, accessToken)
        }
        r.Header.Set("nombre","valor")
        fmt.Println(r.Header)
        
        fmt.Fprintf(w, "\nHola Mundo")
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Header().Add("nombre","valor del header")
        values := r.URL.Query()
        values.Add("Name","Daniel")
        values.Add("Course","Go Web")
        values.Add("Job","Freelancer")
        r.URL.RawQuery = values.Encode()
        fmt.Println(r.URL)
        fmt.Fprintf(w, "Hola Mundo")
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Header().Add("nombre","valor del header")
        fmt.Println(r.URL.Query())
        name := r.URL.Query().Get("name")
        if len(name) != 0{
            fmt.Println(name)
        }
        fmt.Fprintf(w, "Hola Mundo")
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("El metodo es " + r.Method)
        switch r.Method {
            case "GET":
                fmt.Fprintf("Metodo Get")
            case "POST":
                fmt.Fprintf("Metodo POST")
            case "PUT":
                fmt.Fprintf("Metodo put")
            case "DELETE":
                fmt.Fprintf("Metodo delete")
            default: 
                http.Error(w, "Metodo no valido", http.StatusBadRequest)
        }
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Header().Add("nombre","valor del header")
        //fmt.Fprintf(w, "Hola Mundo")
        http.NotFound(w,r)
    })
    http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request){
        //fmt.Fprintf(w, "Hola Mundo dos")
        http.Redirect(w, r, "https://www.google.com", 301)
    })
    http.HandleFunc("/tres", func(w http.ResponseWriter, r *http.Request){
        http.Error(w, "Este es un error", http.StatusConflict)
    })
*/
    
    log.Fatal(server.ListenAndServe())
}