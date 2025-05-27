package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			ImprimeLogs()
			
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

	// sites := []string {
	//  "https://www.google.com",
    // 	"https://www.wikipedia.org",	
    // 	"https://www.youtube.com",
    // 	"https://stackoverflow.com",
    // 	"https://github.com",
	// }

	sites := LerSitesDoArquivo()

	for i := 0; i< moitoriamentos; i++{

		for i, site := range sites{
			fmt.Printf("Testando o site: %d %s\n", i, site)
			TestaSite(site)	
		}	
		time.Sleep(delayMonitoramento *time.Second)
		fmt.Println("")
	}
}

func TestaSite(site string){

	resp, err := http.Get(site)	

	if err != nil {
		fmt.Printf("Ocorreu um erro ao acessar o site %s: %v\n", site, err)
		return
	}	
	if resp.StatusCode == 200{
		fmt.Printf("Site: %s, foi carregado com sucesso\n", site)
		RegistraLog(site, true)
	}else {
		fmt.Printf("Site: %s, está com problemas. Status: %d \n", site, resp.StatusCode)
		RegistraLog(site, false)
	}
}

func LerSitesDoArquivo() []string{
	var site []string
	
	//arquivo, err := ioutil.ReadFile("sites.txt")

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Printf("Ocorreu um erro ao abrir o arquivo: %v\n", err)
		return site
	}
	
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		site = append(site, linha)

		if err == io.EOF {
			break	
		}
	}
	arquivo.Close()	
	return site
}	

func RegistraLog(site string, status bool) {
 	
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Ocorreu um erro ao abrir o arquivo de log: %v\n", err)
		return
	}

	// arquivo.Seek(0, io.SeekEnd) // Move o ponteiro do arquivo para o final
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05" ) + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	
	arquivo.Close() 
}

func ImprimeLogs() {
	println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Printf("Ocorreu um erro ao abrir o arquivo de log: %v\n", err)
		return
	}

	fmt.Println(string(arquivo))	
}