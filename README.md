Semantics3 api client in golang. 

### Example
<pre>
package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "github.com/pirsquare/semantics3-go"
)

func main() {
    client := semantics3.NewClient("XXXXXXXX", "XXXXXXXX", "products")
    client.AddParams(map[string]interface{}{"upc": uint64(636926047593)})

    response, err := client.Get()
    if err != nil{
        log.Fatalln(err)
    }

    defer response.Body.Close()

    bits, err := ioutil.ReadAll(response.Body)
    if err != nil{
        log.Fatalln(err)
    }

    fmt.Println(string(bits))

}
</pre>