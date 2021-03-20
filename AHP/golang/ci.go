package main

import "strconv"

type CiInfo struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

func Ci(l LamdaMaks, n float64) CiInfo {

	lStr := strconv.FormatFloat(l.Total.Value, 'f', 3, 64)
	nStr := strconv.FormatFloat(n, 'f', 3, 64)

	result := CiInfo{
		Value: (((l.Total.Value) - (n)) / ((n) - 1)),
		Desc:  "( " + lStr + " - " + nStr + " ) / (" + nStr + " - 1 )",
	}

	return result
}
