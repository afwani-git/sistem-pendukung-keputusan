package main

import (
	"strconv"
)

type Ri float64

type CrInfo struct {
	Value  float64 `json:"value"`
	Desc   string  `json:"description"`
	Consis bool    `json:"consistance"`
}

//TODO: otomatisasi ordo
// info
// Ordo matriks 	1 	2 	3 		4 		5 		6 		7 		8 		9 		10
//Ratio index 		0 	0 	0.58 	0.9 	1.12 	1.24 	1.32 	1.41 	1.46 	1.49
func Cr(c CiInfo, r Ri) CrInfo {

	val := c.Value / float64(r)

	result := CrInfo{
		Value:  val,
		Desc:   "CR = " + strconv.FormatFloat(float64(c.Value), 'f', 3, 64) + " / " + strconv.FormatFloat(float64(r), 'f', 3, 64),
		Consis: val < 0.1,
	}

	return result
}
