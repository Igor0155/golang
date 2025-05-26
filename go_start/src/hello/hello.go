package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	// "reflect"
)

const moitoriamentos = 3
const delayMonitoramento = 3

func main() {

	ExibeIntroducao()

	for {
		ExibeMenu()
		comando := LerComando()
	
		// if comando == 1 {
		// 	println("Monitorando...")
		// }else if comando == 2{
		// 	println("Exibindo Logs...")
		// }else if comando == 0{
		// 	println("Saindo do programa")
		// }else {
		// 	println("Comando não encontrado")
		// }
	
		switch comando {
		case 1:
			IniciarMonitoramento()
		case 2:
			println("Exibindo Logs...")
		case 0:
			println("Saindo do programa")
			os.Exit(0)
		default:
			println("Comando não encontrado")
			os.Exit(-1)
		}
	}
}

func ExibeMenu(){
	fmt.Println("1 - Iniciar o Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func LerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)
	return comandoLido
}


func ExibeIntroducao(){
	// var name string = "Igor"
	// var numFloat float32 = 22.2
	// var idade int = 22
	// fmt.Printf("\no tipo da váriavel Nome é %s", reflect.TypeOf(name))
	// fmt.Printf("\no tipo da váriavel numFloat é %s", reflect.TypeOf(numFloat))
	// fmt.Printf("\no tipo da váriavel idade é %s", reflect.TypeOf(idade))
	// idade := 22

	name := "Igor G."
	version := 1.0
	fmt.Printf("Olá %s, \nEste programa está na versão %0.1f\n", name, version)
}

func IniciarMonitoramento(){
	println("Monitorando...")

	sites := []string {
	 	"https://www.google.com",
    	"https://www.wikipedia.org",	
    	"https://www.youtube.com",
    	"https://stackoverflow.com",
    	"https://github.com",
	}

	for i := 0; i< moitoriamentos; i++{

		for _, site := range sites{
			TestaSite(site)	
		}	
		time.Sleep(delayMonitoramento *time.Second)
	}
}

func TestaSite(site string){

	resp, _ := http.Get(site)	
	if resp.StatusCode == 200{
		fmt.Printf("Site: %s, foi carregado com sucesso\n", site)
	}else {
		fmt.Printf("Site: %s, está com problemas. Status: %d \n", site, resp.StatusCode)
	}
}