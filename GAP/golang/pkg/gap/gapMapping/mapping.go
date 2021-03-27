package gapMapping

import (
	"strconv"

	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
)

// Gap= 𝑀𝑖𝑛𝑖𝑚𝑎𝑙𝑃𝑟𝑜𝑓𝑖𝑙𝑒−𝑇𝑒𝑠𝑡𝑑𝑎𝑡𝑎𝑝𝑟𝑜𝑓𝑖𝑙𝑒
func MappingGap(g gap.Gap) TableGapMapping {

	tableGapMapping := make(TableGapMapping)
	tableAspec := g.AspectTable

	for aspce, val := range g.CadidateTable {
		tableGapMapping[aspce] = make(map[gap.CadidateName]GapMappingInfo)
		for candidateName, val2 := range val {
			tableGapMapping[aspce][candidateName] = GapMappingInfo{
				Name:     candidateName,
				SubAspec: make(map[gap.SubAspec]GapMappingVal),
			}
			for subAspec, val3 := range val2.Value {

				sMinProfile := strconv.Itoa(val3.Value)
				sTesProfile := strconv.Itoa(tableAspec[aspce].SubAspect[subAspec].TargetValue)
				sResultW := strconv.Itoa(val3.Value - tableAspec[aspce].SubAspect[subAspec].TargetValue)
				wValue := g.Wtable[sResultW]

				tableGapMapping[aspce][candidateName].SubAspec[subAspec] = GapMappingVal{
					MapGapValue: val3.Value - tableAspec[aspce].SubAspect[subAspec].TargetValue,
					WeightValue: WeightValue{
						Value:       wValue.Value,
						Description: wValue.Description,
					},
					MinProfile:  val3.Value,
					DatProfile:  tableAspec[aspce].SubAspect[subAspec].TargetValue,
					Description: "( " + sMinProfile + " - " + sTesProfile + ") = " + sResultW,
					FactorType:  tableAspec[aspce].SubAspect[subAspec].Type,
				}
			}
		}
	}
	return tableGapMapping
}
