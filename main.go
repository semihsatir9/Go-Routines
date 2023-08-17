package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com",
		"http://facebook.com",
		"http://sahibinden.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c{
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link,c)
		}(l)
		

	}
	

}

func checkLink(link string,c chan string) {
	_,err := http.Get(link)

	fmt.Println("Started working on ", link)
	if(err != nil){
		fmt.Println(link," is down! ",err)
		c <- link
	} else{
		fmt.Println(link," is responding")
		c <- link
	}

}