package factory

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/calculate"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/gapMapping"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/ranking"
)

type SubAspecVal int

type FCandidateTable map[gap.Aspec]map[gap.CadidateName]map[gap.SubAspec]SubAspecVal

type ResultGap struct {
	GapSpec      gap.SpecTable
	GapMapping   gapMapping.TableGapMapping
	GapCalculate calculate.TableTotalCalcFactory
	GapRanking   ranking.RankingTable
}
