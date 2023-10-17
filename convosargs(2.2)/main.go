package main

//2.2
import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Print(t, " ft = ", t/3.281, " m , ", t, " m = ", t*3.281, " ft", "\n")
		fmt.Print(t, " kg = ", t*2.205, " lb, ", t, " lb = ", t/2.205, " kg", "\n")
	}
}
