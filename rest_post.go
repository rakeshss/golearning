package main

import(
	"bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	jsonData := map[string]string {"firstname" : "Rakesh", "lastname" : "SS"}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println("Some error occured while making POST call!!")
	}else {
		data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
	}
}
