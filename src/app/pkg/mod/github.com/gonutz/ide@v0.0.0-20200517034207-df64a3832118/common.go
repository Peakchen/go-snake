package main

import "errors"

func makeErr(context string, err error) error {
	return errors.New(context + ": " + err.Error())
}

func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}
