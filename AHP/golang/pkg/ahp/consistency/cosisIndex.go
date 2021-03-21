package consistency

import (
	"strconv"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func Ci(l ahp.LamdaMaks, n float64) ahp.CiInfo {

	lStr := strconv.FormatFloat(l.Total.Value, 'f', 3, 64)
	nStr := strconv.FormatFloat(n, 'f', 3, 64)

	result := ahp.CiInfo{
		Value: (((l.Total.Value) - (n)) / ((n) - 1)),
		Desc:  "( " + lStr + " - " + nStr + " ) / (" + nStr + " - 1 )",
	}

	return result
}
