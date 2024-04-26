package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for strIndex, strLine := range strings {
		floatPrice, err := strconv.ParseFloat(strLine, 64)
		if err != nil {
			// fmt.Println("Floating price failed")
			fmt.Println(err)
			return nil, errors.New("Floating price failed")
		}
		floats[strIndex] = floatPrice
	}
	return floats, nil
}