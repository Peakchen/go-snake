package common

import (
	"math"
	//"fmt"
)

func realrnd(seed int) float64 {
	seedtmp := (seed*9301 + 49297) % 233280 //为何使用这三个数: https://www.zhihu.com/question/22818104
	return float64(seedtmp) / float64(233280.0)
}

func RandInt(number, seed int) int {
	originRand := math.Ceil(realrnd(seed) * float64(number))
	return int(originRand)
}

func RandOne(seed int) float64 {
	return math.Ceil(realrnd(seed)*float64(9999)) / float64(10000)
}

func RandFloat(number float64, seed int) float64 {
	if number == 0.0 {
		return number
	}
	return math.Ceil(realrnd(seed)*(number-1)) / number
}

func RandFloat2Max(number float64, seed int) float64 {
	if number == 0.0 {
		return number
	}
	return math.Ceil(realrnd(seed)*(number)) / number
}
