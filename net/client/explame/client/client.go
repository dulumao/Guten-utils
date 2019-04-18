package client

import (
	"fmt"
	"github.com/dulumao/Guten-utils/net/client"
)

func main() {
	c := client.Default()

	// c.NewEndPoint("http://114.115.143.245:20181")
	// c.SetBasicAuth("test", "test1234")
	// // c.SetMethod("getinfo", map[string]string{
	// // 	"a": "a1",
	// // 	"b": "b1",
	// // 	"c": "c1",
	// // 	"d": "d1",
	// // })
	// c.SetMethod("getinfo")
	c.NewEndPoint("http://180.76.239.140:8512")
	// c.SetBasicAuth("test", "test1234")
	// c.SetMethod("getinfo", map[string]string{
	// 	"a": "a1",
	// 	"b": "b1",
	// 	"c": "c1",
	// 	"d": "d1",
	// })
	c.SetMethod("web3_clientVersion")

	if r, err := c.GetResponse(); err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v \n", r)
	}

	if s, err := c.GetBody(); err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v \n", s)
	}

	r := &struct {
		Jsonrpc string `json:jsonrpc"`
		Id      int    `json:id"`
		Result  string `json:result"`
	}{}

	c.GetJson(r)

	fmt.Printf("%+v \n", r)
}
