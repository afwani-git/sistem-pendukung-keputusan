package globalPriority

import "github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"

func joinAllEign(a *[]Alternative) ValueTable {

	vTable := make(ValueTable)

	for _, val := range *a {

		c := ahp.Col{}

		for key2, val2 := range val.Values.Avg {
			c[key2] = val2.Value
		}

		vTable[val.Key] = c

	}

	return vTable
}
