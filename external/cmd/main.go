package main

import (
	"flag"
	"fmt"
	"ipchecker/internal"
)

func main() {
	ip := flag.String("subnet", "", "Subnet que pertende pesquisar")

	flag.Parse()

	fmt.Println(internal.Check(*ip))
}
