package main

import (
	"fmt"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/board/moebooru"
)

func main() {
	board := moebooru.YandeReConfig
	board.BuildAuth("xxx", "xxx")

	board.Query = getmoe.Query{
		Tags: []string{"nipples", "rating:e"},
		Page: 1,
	}

	posts, err := board.RequestAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range posts {
		fmt.Println(p.FileURL)
	}
}
