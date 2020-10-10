package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	for _, ip := range os.Args[1:]{
		api := "http://freeapi.ipip.net/"+ip
		resp, err := http.Get(api)
		if err != nil {
			panic(err)
			os.Exit(1)
		}
		fmt.Printf("http status -> %s \n", resp.Status)
		header := resp.Header
		for k,v := range header{
			fmt.Printf("%s -> %s \n", k,  v)
		}
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil{
			panic(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)

		mulit(ip)
	}


}

func mulit(ip string) {
	ch := make(chan []byte)
	for i:= 0; i < 10; i ++ {
		go fetch(ip, ch)
	}
	for j := 0; j < 10; j ++ {
		b := <- ch
		fmt.Printf("%s\n", b)
	}
}

func fetch(ip string, ch chan <- []byte)  {
	api := "http://freeapi.ipip.net/"+ip
	resp, err := http.Get(api)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		panic(err)
		os.Exit(1)
	}
	ch <- b
}

