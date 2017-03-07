package viewmodels

import (
	"net/http"
	"log"
	"strings"
	"io/ioutil"
)

var texto string

func PostModelCola(w http.ResponseWriter, req *http.Request) bool {
    x := LeerGuardarPila(w,req)
    ap := ArregloPilaC(x)
    y := GuardarCola()
    ac := ArregloCola(y)
    z:=compararPilaCola(ap,ac)
    return z
}

func LeerGuardarPila(rw http.ResponseWriter, req *http.Request) *Pila{
    log.Println("este es un flag postpilacola")
    var pila *Pila = NewPila()
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    texto = string(body)
    arreglo := strings.Split(texto,"")
    for i := 0; i < len(arreglo); i++{
        if arreglo[i] != " "{
            pila.Push(arreglo[i])
        }
    }
    return pila
}
func NewPila() *Pila {
    var pila *Pila = new(Pila)
    pila.profundidad = 0
    return pila
}

func ArregloPilaC(p *Pila)[]string{
    var profundidad int
    profundidad = int(p.profundidad)
    arregloPila := make([]string, profundidad)  
    for i:=0; i< profundidad ; i++{
        arregloPila[i]= p.Pop().(string)
    }   
    return arregloPila
}

func GuardarCola() *Cola{
    var cola *Cola = NewCola()
    arreglo := strings.Split(texto,"")
    for i := 0; i< len(arreglo); i++{
        if arreglo[i] != " "{
            cola.Agregar(arreglo[i])
        }
    }
    return cola
}

func NewCola() *Cola {
    cola := &Cola{}
    cola.cola = make([]interface{}, 0)
    cola.largo = 0
    return cola
}

func (q *Cola) Agregar(value interface{}) {
    q.largo += 1
    q.cola = append(q.cola, value)
}

func ArregloCola(c *Cola) []string{
    var largo int
    largo = c.Largo()
    arregloCola := make ([]string,largo)
    for i:=0 ; i < largo ; i++ {
        arregloCola[i]=c.Quitar().(string)
    }
    return arregloCola
}

func (cola *Cola) Largo() int {
    return cola.largo
}

func (q *Cola) Quitar() interface{} {
    tmp := q.cola[0]
    q.cola = q.cola[1:]
    q.largo -= 1
    return tmp
}

func compararPilaCola(x []string,y []string) bool{
    for i := 0; i < len(x); i ++ {
        if x[i] != y[i]{
            return false
        }
    }
    return true
}

/*
func LeerGuardarCola(rw http.ResponseWriter, req *http.Request) *Pila{
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

}*/
