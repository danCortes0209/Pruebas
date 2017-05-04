/*
Tipos Compuestos

parte 1

- se declara un array de 5 cadenas cada uno, los inicializaremos con 0
- se declara otro arreglo de 5 strings, los inicializaremos con los valores de strings
- asignar el segundo arreglo al primero y desplegar los resultados del primero en ese arreglo
- mostrar el valor de la cadena y direccion del arreglo

parte 2

- declarar un slice de enteros vacios
- crear un loop que meta 10 valores a nuestro slice
- iterar sobre el slice y mostrar cada valor
- declarar un slice de 5 cadenas strings e inicializar dicho slice con valores string
- desplegar o imprimir todos los elementos de nuestro slice
- tomar un slice del primer y segundo indice
- desplegar la posicion y el valor de cada uno de estos elementos en el nuevo slice

parte 3
 
- Declarar un mapa de valores enteros con un string como llave
- LLenar el mapa con 5 valores
- iterar sobre el mapa para mostrar los pares llave/valor

Extra: Casos de uso de range

La funcion "range" itera por los elementos de la mayoria de las estructuras de datos, muy util para los mapas, itera sobre el rango de alcance de la estructura de datos.
Similar a la funcion .lenght() de otros lenguajes
*/

package main

import(
    "fmt"
)

func main(){
    arrayD()
    sliceD()
    mapD()
    rangeD()
}

func arrayD(){
    //declarar arreglos de string que guarden los nombres
    var nombres[5]string
    
    //declarar un arreglo pre-inicializado con valores string
    amigos:= [5]string{"Raquel","Adriana","Sebastian","Vanesa","Elizabeth"}
    
    //asignar el arreglo de los amigos al arreglo nombres
    nombres = amigos
    
    //Mostrar cada uno de los valores string y la direccion de cada uno del segundo arreglo
    for i, nombre := range nombres{
        fmt.Println(nombre, &nombres[i])
    }
}

func sliceD(){
    //slice de enteros vacios
    var numeros []int
    
    //meter 10 valores al slice
    for i:=0; i<10; i++{
        numeros = append(numeros, i*10)
    }
    
    //mostrar cada valor
    for _, numero := range numeros {
        fmt.Println(numero)
    }
    
    //declarar slice de strings
    frutas := []string{"manzana","naranja","pera","sandia","aguacate"}
    
    //mostrar cada indice o posicion y cada nombre
    for i, fruta := range frutas {
        fmt.Printf("INDEX: %d Fruta: %s\n", i, fruta)
    }
    
    //tomar un slice de indice uno y dos
    slice := frutas[4:5]
    
    //mostrar el valor del nuevo slice
    for i, fruta := range slice {
        fmt.Printf("Index %d Fruta: %s\n", i, fruta)
    }
}

func mapD(){
    //Declarar y hacer el mapa
    departamentos := make(map[string]int)
    
    //Inicializar los datos en el mapa
    departamentos["Devs"] = 25
    departamentos["Marketing"] = 50
    departamentos["Ejecutivos"] = 4
    departamentos["Ventas"] = 60
    departamentos["Mantenimiento"] = 8
    
    //Desplegar por medio de iteracion el valor de cada par llave/valor
    for key, value := range departamentos {
        fmt.Printf("Dept: %s Personal: %d\n", key, value)
    }
}

func rangeD(){
    numeros := []int{2,4,6}
    suma := 0
    for _, num := range numeros {
        suma += num
    }
    
    fmt.Println("Suma: ", suma)
    
    for i, num := range numeros {
        if num == 3 {
            fmt.Println("Index: ", i)
        }
    }
    algo := map[string]string {"a":"auto","b":"bebé"}
     for key, value := range algo {
        fmt.Printf("Letra: %s Ejemplo: %s\n", key, value)
    }
}


