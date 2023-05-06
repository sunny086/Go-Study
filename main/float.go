package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func FloatCompare(f1, f2 interface{}) (n int, err error) {
	var f1Dec, f2Dec decimal.Decimal
	switch f1.(type) {
	case float64:
		f1Dec = decimal.NewFromFloat(f1.(float64))
		switch f2.(type) {
		case float64:
			f2Dec = decimal.NewFromFloat(f2.(float64))
		case string:
			f2Dec, err = decimal.NewFromString(f2.(string))
			if err != nil {
				return 2, err
			}
		default:
			return 2, errors.New("FloatCompare() expecting to receive float64 or string -- 0")
		}
	case string:
		f1Dec, err = decimal.NewFromString(f1.(string))
		if err != nil {
			return 2, err
		}
		switch f2.(type) {
		case float64:
			f2Dec = decimal.NewFromFloat(f2.(float64))
		case string:
			f2Dec, err = decimal.NewFromString(f2.(string))
			if err != nil {
				return 2, err
			}
		default:
			return 2, errors.New("FloatCompare() expecting to receive float64 or string -- 1")
		}
	default:
		return 2, errors.New("FloatCompare() expecting to receive float64 or string -- 2")
	}
	return f1Dec.Cmp(f2Dec), nil
}

func main() {
	a := 70
	b := "70"
	fmt.Println(FloatCompare(a, b))

	max := math.Max(10.34, 70)
	fmt.Println(max)
}
