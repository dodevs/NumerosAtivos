package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var numberReg = regexp.MustCompile(`((55)(\d\d)(9\d+))@c\.us`)

func presence2number(presence Presence) Number {
	groups := numberReg.FindStringSubmatch(presence.ID)
	country, _ := strconv.Atoi(groups[2])
	ddd, _ := strconv.Atoi(groups[3])
	number := groups[1]

	return Number{
		country:  country,
		ddd:      ddd,
		number:   number,
		valid:    true,
		lastView: presence.T,
	}
}

func errorHandler(text string, err error) {
	if err != nil {
		fmt.Println(text, err)
	}
}
