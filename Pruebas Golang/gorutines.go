//Go routines
package main
import (
    "fmt"
    "runtime"
    "sync"
)
//init se va a ejecutar antes de main
func init (){
    //Alojando o reservando (allocate) un procesador logico para que lo use el scheduler.
    runtime.GOMAXPROCS(1)
}

func main(){
    //declarar un grupo de espera para iniciar el conteo
    var wg sync.WaitGroup
    wg.Add(2)
    fmt.Println("Iniciar nuestras goroutines")
    
    //funcion anonima y crear una goroutine a partir de ahi
    
    go func(){
        //cuenta regresiva del 100 al 0
        for count := 100; count>=0; count-- {
            fmt.Printf("[A:%d]\n", count)   
        }
        wg.Done()
    }()
    
    //declarar una funcion anonima y crear una gorutine
    go func(){
        for count := 0; count<100; count++ {
            fmt.Printf("[B:%d]\n", count)
        }
        wg.Done()
    }()
    
    
    //esperar a que terminen
    fmt.Println("Waiting...")
    wg.Wait()
    
    fmt.Println("Finalizando el programa")
}































