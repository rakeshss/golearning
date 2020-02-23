package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	fmt.Println("starting application")
	response, err  := http.Get("https://httpbin.org/ip") 

	if err != nil {
		fmt.Println("The HTTP request failed with error!!")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

