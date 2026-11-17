package main

import (
	"fmt"
)

func main() {

	events, err := LatestPlayClient()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range events {
		fmt.Println(event.LastPlay.Text)
	}
}
