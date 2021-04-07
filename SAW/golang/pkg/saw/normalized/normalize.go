package normalized

import (
	"sort"
	"strconv"

	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw"
)

// Rij = ( Xij/max{Xij} ) benefit

// Rij = ( min{Xij} / Xij) cost

type formula map[saw.CriteriaName]map[saw.Attr][]float64

func checkAttr(at saw.AlternativeTable) formula {

	formula := make(formula)

	// counter := 0

	for _, key := range at {
		for cName, key2 := range key.CriteriaTable {
			if formula[cName] == nil {
				formula[cName] = make(map[saw.Attr][]float64)
			}
			formula[cName][key2.Criteria.Attr] = append(formula[cName][key2.Criteria.Attr], key2.Values)
			formula[cName][key2.Criteria.Attr] = sort.Float64Slice(formula[cName][key2.Criteria.Attr])

			sort.Slice(formula[cName][key2.Criteria.Attr], func(i, j int) bool {
				return formula[cName][key2.Criteria.Attr][i] < formula[cName][key2.Criteria.Attr][j]
			})
		}
	}

	return formula
}

func NormalizedValue(ct saw.CriteriaTable, at saw.AlternativeTable) NormalizedTable {

	normTable := make(NormalizedTable)

	formula := checkAttr(at)

	for cName, _ := range ct {
		normTable[cName] = NormalizedInfo{
			CriteriaName:    cName,
			AlternativeList: make(AlternativeList),
		}
		for aName, key2 := range at {

			attrType := key2.CriteriaTable[cName].Criteria.Attr

			if attrType == saw.BENEFIT {
				normTable[cName].AlternativeList[aName] = NormalizedAltInfo{
					Values: key2.CriteriaTable[cName].Values / formula[cName][attrType][len(formula[cName][attrType])-1],
					Desc:   strconv.FormatFloat(key2.CriteriaTable[cName].Values, 'f', 2, 64) + " / " + strconv.FormatFloat(formula[cName][attrType][len(formula[cName][attrType])-1], 'f', 2, 64),
				}
			} else {
				normTable[cName].AlternativeList[aName] = NormalizedAltInfo{
					Values: formula[cName][attrType][0] / key2.CriteriaTable[cName].Values,
					Desc:   strconv.FormatFloat(formula[cName][attrType][0], 'f', 3, 64) + " / " + strconv.FormatFloat(key2.CriteriaTable[cName].Values, 'f', 3, 64),
				}
			}

		}
	}

	return normTable
}
