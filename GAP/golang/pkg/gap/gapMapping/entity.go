package gapMapping

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
)

type WeightValue struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type GapMappingVal struct {
	MapGapValue int            `json:"mapGapValue"`
	WeightValue WeightValue    `json:"weightValue"`
	MinProfile  int            `json:"minProfile"`
	DatProfile  int            `json:"datProfile"`
	Description string         `json:"description"`
	FactorType  gap.FactorType `json:"factorType"`
}

type GapMappingInfo struct {
	Name     gap.CadidateName               `json:"name"`
	SubAspec map[gap.SubAspec]GapMappingVal `json:"subAspec"`
}

type TableGapMapping map[gap.Aspec]map[gap.CadidateName]GapMappingInfo
