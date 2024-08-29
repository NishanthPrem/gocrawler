package main

import (
	"fmt"
	"os"
)

func main() {
	parseArgs()

}

func parseArgs() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Println("starting crawl of:", args[0])
	}
}
