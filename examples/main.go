package main

import (
	"fmt"
	"github.com/pirsquare/semantics3-golang"
	"io/ioutil"
)

func main() {
	client := semantics3.NewClient("XXXXXXXXXX", "XXXXXXXXXX", "products")
	client.AddParams(map[string]interface{}{"upc": uint64(636926047593)})

	response, err := client.Get()
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bits))

}
