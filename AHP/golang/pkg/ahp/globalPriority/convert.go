package globalPriority

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func convert(ce ahp.Ahp, v ValueTable, c ahp.CandidateTable) ConvertedTable {

	covTable := make(ConvertedTable)

	for cName, val := range c {
		covTable[cName] = make(map[ahp.AspecName]float64)
		for aspec, val2 := range val {
			result := v[string(aspec)][string(val2)] * ce.Avg[string(aspec)].Value
			covTable[cName][aspec] = result
			covTable[cName]["jumlah"] += result
		}
	}

	return covTable
}
