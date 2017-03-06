package viewmodels

type ViewModel struct {
    Title string
    SignedIn bool
    Active string
}
type MyModel struct {
	Arreglo string
	SignedIn bool
    Result_data string

}

type ObjetoPila struct {
    item interface{}
    next *ObjetoPila
}
// La pila es una estructura base para LIFO (Last in First out)

type Pila struct {
    op    *ObjetoPila
    profundidad uint64
}