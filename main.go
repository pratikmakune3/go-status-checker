package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Slice of strings
	links := []string{
		"https://platform.excessmaterialsexchange.com/",
		"https://test.excessmaterialsexchange.com/",
		"https://demo.excessmaterialsexchange.com/",
		"https://testmp.excessmaterialsexchange.com/",
		"https://marketplace.excessmaterialsexchange.com/",
		"https://testwear2go.circular.exchange/",
		"https://protrade.excessmaterialsexchange.com/",
		// Imagine we have 1000 of links... blocking code is BAD!
	}

	c := make(chan string)

	for _, link := range links {
		// We can use go keyword to create a goroutine
		// only infront of a func call
		go checkLink(link, c)
	}


	// <- c blocking code, main routine is going to be blocking the code
	// fmt.Println(<- c) 
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	
	// fmt.Println(<- c)

	for {
		// Value that coming from <- c is blocking call 
		// so, let's create a go routine out of this blocking call
		// fmt.Println(<- c)

		// sleep CURRENT go-routine -> Main routine -> BAD idea
		// time.Sleep(5 * time.Second)
		go checkLink(<- c, c)
	}

}

// Serial execution...
func checkLink(link string, c chan string) {

	time.Sleep(10 * time.Second)

	// Returns response, err
	_, err := http.Get(link) // blocking call

	if err != nil {
		fmt.Println(link, "might be down!!!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up :)")
	// c <- "Yep running"
	c <- link
}