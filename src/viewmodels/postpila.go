package viewmodels

import (
	"net/http"
	"log"
	"strings"
	"io/ioutil"
)
func PostModel(w http.ResponseWriter, req *http.Request) []string {
    x := LeerGuardar(w,req)
    y := ArregloPila(x)
    return y
}


func LeerGuardar(rw http.ResponseWriter, req *http.Request) *Pila{
    log.Println("este es un flag")
    var pila *Pila = New()
        body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    s := string(body)
    log.Println(`DANNNIII`)
    log.Println(s)
    
    arreglo := strings.Split(s,"")
    for i := 0; i < len(arreglo); i++{
        pila.Push(arreglo[i])
    }
    //return 
    log.Println("este es un flag 3 ")
    //rw.Write([]byte("Hello VARIABLE!"))
    //rw.Write(y)
    return pila

}

func New() *Pila {
    var pila *Pila = new(Pila)
    pila.profundidad = 0
    return pila
}

// Ingresar un item dado a la pila
func (pila *Pila) Push(item interface{}) {
    pila.op = &ObjetoPila{item: item, next: pila.op}
    pila.profundidad++
}

// Elimina el tope de la pila y lo retorna

func (pila *Pila) Pop() interface{} {
    if pila.profundidad > 0 {
        item := pila.op.item
        pila.op = pila.op.next
        pila.profundidad--
        return item
    }
    return nil
}

// Retorna el tope de la pila sin eliminarlo

func (pila *Pila) Peek() interface{} {
    if pila.profundidad > 0 {
        return pila.op.item
    }
    return nil
}


func ArregloPila(p *Pila)[]string{
    var profundidad int
    profundidad = int(p.profundidad)
    arregloPila := make([]string, profundidad)  
    for i:=0; i< profundidad ; i++{
        arregloPila[i]= p.Pop().(string)
    }   
    return arregloPila
}