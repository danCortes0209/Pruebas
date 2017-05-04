//hello world

package main

import (
    "fmt"
)
func main(){
    fmt.Println("Hello World :D")
    
    //mandamos llamar un metodo o funcion
    whatrudoing()
    
    //nuevo metodo de variables :D
    newVariablesTry()
    
    //nuevo metodo para el ciclo for
    cicloFor()
}

func whatrudoing(){
    //queremos imprimir la cadena golang
    fmt.Println("go" + "lang")
    
    //queremos aprender a sumar/restar enteros
    fmt.Println("5+5 =", 5+5)
    
    //queremos aprender a manejar los flotantes
    fmt.Println("7.5/3.2 =", 7.5/3.2)
    //OJO, NO HAY CONVERSION IMPLICITA, TIENES QUE DECLARARLA COMO EXPLICITA
    
    
    //Manejo de booleanos
    fmt.Println("Resultados =")
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!false)
}

func newVariablesTry(){
    //declarando una sola variable
    var a string = "Inicial"
    fmt.Println(a)
    
    //declarando mas de una variable
    var b, c int = 4, 5
    fmt.Println(b, c)
    
    //declarando variables con resultado booleano
    var d = true
    fmt.Println(d)
    
    //variables no inicializadas
    var e int 
    fmt.Println(e)
    
    //otra forma de declarar la variable
    f:= "short"
    fmt.Println(f)
}

func cicloFor(){
    alumnos := 1
    for alumnos <= 3{
        fmt.Println(alumnos)
        alumnos = alumnos + 1
    }
    
    for calificaciones := 7; calificaciones <=9; calificaciones++{
        fmt.Println(calificaciones)
    }
}
