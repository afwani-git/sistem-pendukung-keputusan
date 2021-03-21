package consistency

import (
	"strconv"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

// info
// Ordo matriks 	1 	2 	3 		4 		5 		6 		7 		8 		9 		10
//Ratio index 		0 	0 	0.58 	0.9 	1.12 	1.24 	1.32 	1.41 	1.46 	1.49

type Ri float64
type Ordo int

func OrdoToRi(o Ordo) Ri {

	ratio := map[int]Ri{
		1:  0,
		2:  0,
		3:  0.58,
		4:  0.9,
		5:  1.12,
		6:  1.24,
		7:  1.32,
		8:  1.41,
		9:  1.46,
		10: 1.49,
	}

	return ratio[int(o)]
}

func Cr(c ahp.CiInfo, r Ri) ahp.CrInfo {

	val := c.Value / float64(r)

	result := ahp.CrInfo{
		Value:  val,
		Desc:   "CR = " + strconv.FormatFloat(float64(c.Value), 'f', 3, 64) + " / " + strconv.FormatFloat(float64(r), 'f', 3, 64),
		Consis: val < 0.1,
	}

	return result
}
