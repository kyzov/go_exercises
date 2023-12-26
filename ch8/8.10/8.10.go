package main

import (
	"fmt"
	"log"
	"os"

	"links"
)

var cancel = make(chan struct{})

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url, cancel)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
