package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oreqizer/go-relaygen/generator"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: <name> <id field>")
		fmt.Println("  examples:")
		fmt.Println("  go-relaygen User LocalID")
		fmt.Println("  go-relaygen User Embedded.ID")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := generator.Generate(os.Args[1], os.Args[2], wd); err != nil {
		log.Fatal(err)
	}
}
