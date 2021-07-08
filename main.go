package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlags = flag.String("p","", "Example: -p <PROJECT_ID>")
	flag.Parse()
	switch {
	case *nFlags == "project1":
		fmt.Println("Project ID: ", *nFlags)
	case *nFlags == "project2":
		fmt.Println("Project ID: ", *nFlags)
	default:
		fmt.Println("Error, use -h for help!")
	}
}
