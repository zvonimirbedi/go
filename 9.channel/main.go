package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"http://amazon.com",
		"http://adasdadasdasdadasd.com",
	}

	// channel of type string
	c := make(chan string)

	for _, link := range links {
		// make go routine for every request
		go checkSite(link, c)
	}

	// wait for all go routines, recieve value from channel
	/*
		for range links {
			// print from main routine
			fmt.Println(<-c)
		}
	*/

	// infinte wait for all go routines, recieve value from channel
	/*
		for {
			go checkSite(<-c, c)
		}
	*/
	// wait fot the channel to return some value, then assign it to 'l' (link value)
	for l := range c {
		// function literals - lambda function
		go func(link string) {
			// pause next go routine (request call) for x seconds
			time.Sleep(time.Second * 5)
			checkSite(link, c)
		}(l)
	}
}

func checkSite(url string, c chan string) {
	_, error := http.Get(url)
	if error != nil {
		// print from go routine
		fmt.Println("Link is DOWN:", url)
		// c <- "Links is down, check later"
		c <- url
	} else {
		// print from go routine
		fmt.Println("Link is GOOD:", url)
		// c <- "Link is OK, continue"
		c <- url
	}
}
