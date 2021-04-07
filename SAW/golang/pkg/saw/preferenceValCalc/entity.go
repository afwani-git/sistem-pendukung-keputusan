package preferenceValCalc

import "github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw"

type PValuInfo struct {
	Name  saw.AlternativeName `json:"name"`
	Value float64             `json:"value"`
	Desc  string              `json:"description"`
}

type PreferenceValTable map[saw.AlternativeName]PValuInfo
