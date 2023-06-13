package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir logs...")
		case 3:
			fmt.Println("Sair...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)

		}
	}

}

func exibeIntroducao() string {
	return "Paula"
}

func exibeMenu() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir logs")
	fmt.Println("3 - sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	var sites [4]string
	sites[0] = "www.google.com.br"
	sites[1] = "www.caelum.com.br"
	sites[2] = "www.arquivei.com.br"

	fmt.Println(sites)

	site := "http://alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}
