package parser

import (
	"strings"
	"errors"
)

var (
	GroupOneErr = errors.New("first time split err.")
	GroupTwoErr = errors.New("second time split err.")
	GroupThreeErr = errors.New("three time split err.")
)

/*
	data format: /router?aaa=111&bbb=222...
*/

func ParseURL(context string)(ret map[string]interface{}, err error){

	ret = make(map[string]interface{})

	g1 := strings.Split(context, "?")
	if g1 == nil || len(g1) <= 1 {
		err = GroupOneErr
		return
	}

	g2 := strings.Split(g1[1], "&")
	if g2 == nil || len(g2) <= 1 {
		err = GroupTwoErr
		return
	}

	for _, item := range g2 {

		iv := strings.Split(item, "=")
		if iv == nil || len(iv) <= 1 {
			err = GroupThreeErr
			return
		}

		ret[iv[0]] = iv[1]

	}

	return
}