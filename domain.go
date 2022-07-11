// package
// date 11/7/2022
// author: prr, azul software
// copyright 2022 prr, azul software
//

package main

import (
    "fmt"
    "os"
//   nchp "github.com/namecheap/go-namecheap-sdk/v2"
    nchp "namchp/nchplib"
)

func main() {

//    url := "https://api.sandbox.namecheap.com/xml.response"

    client := nchp.NewClient(&nchp.ClientOptions{
        UserName:   "priemen02",
        ApiUser:    "priemen02",
        ApiKey: "2ae0427ef4344647b43de5ceb6a5c9ea",
        ClientIp: "217.69.0.247",
        UseSandbox: true,
    })

    _, err := client.Domains.GetList(&nchp.DomainsGetListArgs{})
    if err != nil {
        fmt.Printf("error get list: %v\n", err)
        os.Exit(-1)
    }

    fmt.Println("**** namecheap api success ****")
}
