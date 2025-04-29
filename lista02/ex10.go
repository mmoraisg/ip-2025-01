package main
import "fmt"

func main() {
    var a, b int
    
    precos:= map[int]map[int]float64 {
        1: {1:500.00, 2:900.00}, 
        2: {1:350.00, 2:650.00},
        3: {1:350.00, 2:600.00},
        4: {1:300.00, 2:550.00},
    }
    
    regioes:= map[int]string {
        1: "Regiao Norte", 
        2: "Regiao Nordeste",
        3: "Regiao Centro-Oeste",
        4: "Regiao Sul",
    }
    
        fmt.Print("ESCOLHA SEU DESTINO.\n 1- Regiao Norte\n 2- Regiao Nordeste\n 3- Regiao Centro-Oeste\n 4- Regiao Sul\n DIGITE O DIGITO RERFERENTE AO DESITNO:")
        fmt.Scan(&a)
        fmt.Print("ESCOLHA O TIPO DE VIAGEM.\n 1- Ida\n 2- Ida e Volta\n DIGITE O DIGITO RERFERENTE AO TIPO DE VIAGEM:") 
        fmt.Scan(&b)
        
        
    if _, destinovalido:= precos[a]; !destinovalido {
        fmt.Print("DESTINO INVALIDO.")
        return
    }
    if _, retornovalido:=precos[a][b]; !retornovalido {
        fmt.Print("TIPO DE VIAGEM INVALIDO.")
        return
    } 
    preco := precos[a][b]
	tipoViagem := "Só Ida"
	if b == 2 {
		tipoViagem = "Ida e Volta"
	}

	fmt.Println("\nResumo da viagem:")
	fmt.Println("Destino:", regioes[a])
	fmt.Println("Tipo de viagem:", tipoViagem)
	fmt.Printf("Preço da passagem: R$ %.2f\n", preco)
}
