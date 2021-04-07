package normalized

import "github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw"

type NormalizedAltInfo struct {
	Values float64 `json:"values"`
	Desc   string  `json:"description"`
}

type AlternativeList map[saw.AlternativeName]NormalizedAltInfo

type NormalizedInfo struct {
	CriteriaName    saw.CriteriaName `json:"createriaName"`
	AlternativeList AlternativeList  `json:"alternativeList"`
}

type NormalizedTable map[saw.CriteriaName]NormalizedInfo
