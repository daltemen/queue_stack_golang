package viewmodels

func GetPilaCola() ViewModel {
    result := ViewModel{
        Title: "PilaCola",
        SignedIn: false,
        Active: "pilacola",
    }
    return result
}