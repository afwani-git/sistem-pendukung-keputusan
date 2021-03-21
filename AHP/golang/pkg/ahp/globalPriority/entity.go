package globalPriority

import "github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"

type Cgp struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

type InfoRank struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
	Key   string  `json:"key"`
}

type GpInfo struct {
	Table   map[string]ahp.Col
	Process map[string]Cgp   `json:"process"`
	Ranking map[int]InfoRank `json:"ranking"`
}

type Alternative struct {
	Desc   string   `json:"description"`
	Key    string   `json:"key"`
	Values *ahp.Ahp `json:"list"`
}
