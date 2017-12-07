package main

import (
		"fmt"
		"time"
	)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	
	c := make(chan string)
	
	for _, link := range links {
		checkLink(link, c)
	}
	
//	for i := 0;i < len(links); i++ {
//		fmt.Println(<-c)	
//	}
//	for {
//		go checkLink(<-c, c)
//	}
	for l := range c {
		checkLink(l, c)
	}


}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Seconds)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
//		c <- "Might be down I think"
		c<-link
		return
	}
	
	fmt.Println(link, "is up!")
//	c <- "Yep its up"
	c<-link
}