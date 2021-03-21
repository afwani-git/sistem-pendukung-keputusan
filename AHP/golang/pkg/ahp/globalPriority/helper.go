package globalPriority

import (
	"strconv"
	"strings"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func extractValues(t *map[string]ahp.Col, a *[]Alternative) {
	for _, val := range *a {

		c := ahp.Col{}

		for key2, val2 := range val.Values.Avg {
			c[key2] = val2.Value
		}

		(*t)[val.Key] = c

	}
}

func calculateValues(c *ahp.Ahp, gpI *GpInfo) map[string]Cgp {

	type Info struct {
		val  float64
		nums []string
	}

	aInfo := map[string][]Info{}

	for key, val := range gpI.Table {
		for key2, val2 := range val {

			v1 := strconv.FormatFloat(val2, 'f', -3, 64)
			v2 := strconv.FormatFloat(c.Avg[key].Value, 'f', -3, 64)

			aInfo[key2] = append(aInfo[key2], Info{
				val:  val2 * c.Avg[key].Value,
				nums: []string{"( " + v1 + " x " + v2 + " )"},
			})
		}
	}

	aInfo2 := map[string]Info{}
	for key, val := range aInfo {
		info := Info{}
		for _, val2 := range val {
			info.val += val2.val
			info.nums = append(info.nums, val2.nums[0])
		}
		aInfo2[key] = info
	}

	process := map[string]Cgp{}
	for key, val := range aInfo2 {
		cgp := Cgp{
			Value: val.val,
			Desc:  strings.Join(val.nums, " + ") + " = " + strconv.FormatFloat(val.val, 'f', 3, 64),
		}
		process[key] = cgp
	}

	return process
}
