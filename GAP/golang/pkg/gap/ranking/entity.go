package ranking

import "github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"

type ProcessValue struct {
	Name        string  `json:"Name"`
	Description string  `json:"description"`
	FinalScore  float64 `json:"finalSecore"`
}

type ProcessInfo map[gap.CadidateName]ProcessValue

type RankingTable struct {
	Process ProcessInfo `json:"result"`
}
