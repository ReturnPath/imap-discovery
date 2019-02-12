package main

import (
	"fmt"

	discovery "github.com/ReturnPath/contextio-discovery/src"
)

func main() {
	email := "aaron@yahoo.com"
	config, err := discovery.DiscoverImapConfig(email)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}
