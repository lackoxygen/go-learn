package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	URL string = "https://www.baidu.com/"
)

func main()  {
	fmt.Println(http.MethodGet)
	fmt.Println(http.MethodHead)
	fmt.Println(http.MethodPost)
	fmt.Println(http.MethodPut)
	fmt.Println(http.MethodPatch)
	fmt.Println(http.MethodDelete)
	fmt.Println(http.MethodConnect)
	fmt.Println(http.MethodOptions)
	fmt.Println(http.MethodTrace)
	var header = make(http.Header)
	header.Set("version", "1.0.0")
	header.Set("token", "123456")
	header.Add("user-agent", "go")
	fmt.Printf("%v\n", header)
	var resp *http.Response
	var err error
	log.Println("----------head------------")
	resp, err = http.Head(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status => %s\n", resp.Status)
	fmt.Printf("ContentLength => %d\n", resp.ContentLength)
	fmt.Printf("ProtoMajor => %d\n", resp.ProtoMajor)
	fmt.Printf("StatusCode => %d\n", resp.StatusCode)
	respHeader := resp.Header
	for k, v := range respHeader{
		fmt.Printf("%s => %s \n", k, v)
	}
	cookies := resp.Cookies()
	for _, e := range cookies {
		fmt.Printf("%s", e)
	}
	log.Println("----------end------------")
	log.Println("----------get------------")
	resp, err = http.Get(URL)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)
	log.Println("----------end------------")

	log.Println("----------post------------")
	http.Post(URL, "application/json", strings.NewReader("token=123456"))
	log.Println("----------end------------")

	http.PostForm(URL, url.Values{"token":{"123456"}})

	defer resp.Body.Close()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))

	fmt.Printf("serve run at 8000")
}
