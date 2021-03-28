package globalPriority

import "github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"

type GpInfo struct {
	Table          ValueTable     `json:"valueTable"`
	ConvertedTable ConvertedTable `json:"convertedTable"`
}

type ValueTable map[string]ahp.Col
type ConvertedTable map[ahp.CandidateName]map[ahp.AspecName]float64

type Alternative struct {
	Desc   string   `json:"description"`
	Key    string   `json:"key"`
	Values *ahp.Ahp `json:"list"`
}
