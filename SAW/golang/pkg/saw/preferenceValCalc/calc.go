package preferenceValCalc

import (
	"strconv"
	"strings"

	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw"
	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw/normalized"
)

func CalcPreference(nt normalized.NormalizedTable, ct saw.CriteriaTable) PreferenceValTable {

	tmp := make(map[saw.AlternativeName]struct {
		Value []float64
		Desc  []string
	})

	pValueTable := make(PreferenceValTable)

	for cName, key := range nt {
		for aName, key2 := range key.AlternativeList {
			tmp[aName] = struct {
				Value []float64
				Desc  []string
			}{
				Value: append(tmp[aName].Value, key2.Values*ct[cName].Value),
				Desc:  append(tmp[aName].Desc, "( "+strconv.FormatFloat(key2.Values, 'f', 3, 64)+" * "+strconv.FormatFloat(ct[cName].Value, 'f', 3, 64)+" )"),
			}
		}
	}

	for aName, key := range tmp {
		result := PValuInfo{
			Name:  aName,
			Value: 0,
			Desc:  "",
		}

		for _, val := range key.Value {
			result.Value += val
		}

		result.Desc = strings.Join(tmp[aName].Desc, " + ")
		pValueTable[aName] = result
	}

	return pValueTable
}
