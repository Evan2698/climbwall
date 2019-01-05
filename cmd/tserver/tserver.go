package main

import (
	"fmt"
	"net"
	"socks5/core"
	"socks5/security"
)

func main() {

	all, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Print("error", err)
		return
	}

	for {
		one, err := all.Accept()
		if err != nil {
			break
		}

		go func() {
			gcm := security.NewEncryptyMethod("")
			one.Write(security.ToBytes(gcm))
			t := core.NewSocksSocket(one, "zhangweihua", gcm)
			defer t.Close()

			buf, err := t.Read()
			fmt.Println("----------------", buf, err)
			t.Write([]byte("zhangweihua123"))

		}()

	}

}
