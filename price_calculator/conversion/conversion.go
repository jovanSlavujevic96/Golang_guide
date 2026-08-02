package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringsVal := range strings {
		floatVal, err := strconv.ParseFloat(stringsVal, 64)

		if err != nil {
			return nil, err
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
