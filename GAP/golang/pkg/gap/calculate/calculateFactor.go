package calculate

import (
	"strconv"
	"strings"

	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/gapMapping"
)

func CalculateFactor(m gapMapping.TableGapMapping, g gap.Gap) TableTotalCalcFactory {

	table := make(TableTotalCalcFactory)
	// aspecTable := g.AspectTable

	for aspec, val := range m {
		table[aspec] = make(map[gap.CadidateName]CalcFactorInfo)
		for cName, val2 := range val {

			coreFactorAr := make([]string, 0)
			var nCoreFactor float64
			secFactorAr := make([]string, 0)
			var nSecFactor float64

			for _, val3 := range val2.SubAspec {

				switch val3.FactorType {
				case gap.CORE_FACTORY:
					coreFactorAr = append(coreFactorAr, strconv.FormatFloat(val3.WeightValue.Value, 'f', -1, 64))
					nCoreFactor += val3.WeightValue.Value
				case gap.SECONDARY_FACTORY:
					secFactorAr = append(secFactorAr, strconv.FormatFloat(val3.WeightValue.Value, 'f', -1, 64))
					nSecFactor += val3.WeightValue.Value
				}

			}

			pCoFactorStr := strconv.FormatFloat(g.AspectTable[aspec].CoreFactory/100, 'f', -1, 64) + " (" + strconv.FormatFloat(g.AspectTable[aspec].CoreFactory, 'f', -1, 64) + "%)"
			pScFactorStr := strconv.FormatFloat(g.AspectTable[aspec].SecFactory/100, 'f', -1, 64) + " (" + strconv.FormatFloat(g.AspectTable[aspec].SecFactory, 'f', -1, 64) + "%)"
			nCoFactorStr := strconv.FormatFloat(nCoreFactor/float64(len(coreFactorAr)), 'f', -1, 64)
			nScFactorStr := strconv.FormatFloat(nSecFactor/float64(len(secFactorAr)), 'f', -1, 64)
			totalValDesc := pCoFactorStr + " * " + nCoFactorStr + " + " + pScFactorStr + " * " + nScFactorStr

			totalAvgCore := nCoreFactor / float64(len(coreFactorAr))
			totalAvgSec := nSecFactor / float64(len(secFactorAr))
			totalValue := ((g.AspectTable[aspec].CoreFactory / 100) * totalAvgCore) + ((g.AspectTable[aspec].SecFactory / 100) * totalAvgSec)

			table[aspec][cName] = CalcFactorInfo{
				ClcCoreFactor: ClcFactor{
					Value:       totalAvgCore,
					Description: "( " + strings.Join(coreFactorAr, " + ") + " )" + " / " + strconv.Itoa((len(coreFactorAr))),
				},
				ClcSecFactor: ClcFactor{
					Value:       totalAvgSec,
					Description: "( " + strings.Join(secFactorAr, " + ") + " )" + " / " + strconv.Itoa((len(secFactorAr))),
				},
				Name: cName,
				TotalValue: ClcFactor{
					Value:       totalValue,
					Description: totalValDesc,
				},
			}
		}
	}

	return table
}
