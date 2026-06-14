package main

import (
	"flag"
	"fmt"
	"ipchecker/internal"
)

func main() {
	count := flag.Int("c", 10, "Quantidade de ips que vao aparecer")
	busy := flag.Bool("b", true, "Mostrar unicamenta os que estão ocupados")
	ip := flag.String("sn", "", "Subnet que pertende pesquisar")
	saveTxt := flag.Bool("txt", false, "Salvar o resultado em .txt")

	flag.Parse()

	result := internal.Check(*ip)

	for _, obj := range result {
		for range *count {
			if *busy && !obj.IsBusy {
				continue
			}
			fmt.Println(obj.Address, " está ocupado.")
		}
	}

	if *saveTxt {
		fmt.Println("Feature ainda em desenvolvimento!")
	}

}
