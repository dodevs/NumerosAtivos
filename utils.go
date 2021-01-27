package main

import (
	"fmt"
)

func errorHandler(text string, err error) {
	if err != nil {
		fmt.Println(text, err)
	}
}
