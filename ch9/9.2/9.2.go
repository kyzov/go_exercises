package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"popcount"
)

func main() {
	var wg sync.WaitGroup

	for _, str := range os.Args[1:] {
		str := str
		wg.Add(1)
		go func() {
			defer wg.Done()
			n, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				log.Print(err)
			} else {
				fmt.Printf("PopCount of %d = %d\n", n, popcount.PopCount(n))
			}
		}()
	}
	wg.Wait()
}
