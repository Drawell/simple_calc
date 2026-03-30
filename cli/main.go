package main

import (
	"flag"
	"fmt"
	"github.com/Drawell/simple_calc/calc"
	"os"
	"time"
)

func main() {
	var expression string
	flag.StringVar(&expression, "expression", "", "Expression to calculate")

	isWaiting := flag.Bool("wait", false, "Emulate long calculating")

	flag.Parse()

	if *isWaiting {
		time.Sleep(time.Duration(len(expression)) * time.Second)
	}

	res, err := calc.Evaluate(expression)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(res)
		os.Exit(0)
	}
}
