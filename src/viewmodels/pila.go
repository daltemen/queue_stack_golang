package viewmodels

func GetPila() ViewModel {
    result := ViewModel{
        Title: "Pila",
        SignedIn: false,
        Active: "pila",
    }
    return result
}