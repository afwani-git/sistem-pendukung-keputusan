package calculate

import "github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"

type ClcFactor struct {
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}

type CalcFactorInfo struct {
	Name          gap.CadidateName `json:"name"`
	ClcCoreFactor ClcFactor        `json:"coreFactory"`
	ClcSecFactor  ClcFactor        `json:"secondaryFactory"`
	TotalValue    ClcFactor        `json:"totalValue"`
}

type TableTotalCalcFactory map[gap.Aspec]map[gap.CadidateName]CalcFactorInfo
