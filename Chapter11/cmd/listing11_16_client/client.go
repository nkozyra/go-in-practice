package main

import "github.com/nkozyra/cheleven"

func main() {
	if err := cheleven.ConnectToGRPC(); err != nil {
		panic(err)
	}
}
