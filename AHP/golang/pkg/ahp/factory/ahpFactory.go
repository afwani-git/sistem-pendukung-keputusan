package factory

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/consistency"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/globalPriority"
)

type AlternativeTable struct {
	Key   string
	Desc  string
	Table ahp.Row
}

type AhpFactory struct {
	TableCriteria    *ahp.Row
	TableAlternative *[]AlternativeTable
	TableCandidate   ahp.CandidateTable
}

func (a *AhpFactory) CreateCeriteria() ahp.Ahp {
	ceriteria := &ahp.Ahp{}

	ceriteria.Table = *a.TableCriteria
	ceriteria.Sum = a.TableCriteria.ProcessSum()
	ceriteria.Eign = a.TableCriteria.ProcessEigen(ceriteria)
	ceriteria.Avg = a.TableCriteria.ProcessAvg(ceriteria)

	//cek cosistency
	ceriteria.Lmaks = consistency.ProcessLamdaMaks(ceriteria)
	ceriteria.Ci = consistency.Ci(ceriteria.Lmaks, float64(ceriteria.Table.LenCol()))

	ordo := consistency.Ordo(a.TableCriteria.LenCol())
	ceriteria.Cr = consistency.Cr(ceriteria.Ci, consistency.OrdoToRi(ordo))

	return *ceriteria
}

func (a *AhpFactory) CreateAlternatives() []globalPriority.Alternative {

	var alternatives []globalPriority.Alternative

	for _, val := range *a.TableAlternative {
		alternative := &ahp.Ahp{}
		alternative.Table = val.Table
		alternative.Sum = val.Table.ProcessSum()
		alternative.Eign = val.Table.ProcessEigen(alternative)
		alternative.Avg = val.Table.ProcessAvg(alternative)

		//cek cosistency
		alternative.Lmaks = consistency.ProcessLamdaMaks(alternative)
		alternative.Ci = consistency.Ci(alternative.Lmaks, float64(3))

		ordo := consistency.Ordo(val.Table.LenCol())
		alternative.Cr = consistency.Cr(alternative.Ci, consistency.OrdoToRi(ordo))

		//append to alternative list
		alternatives = append(alternatives, globalPriority.Alternative{
			Desc:   val.Desc,
			Key:    val.Key,
			Values: alternative,
		})
	}

	return alternatives
}

func (a *AhpFactory) Create() globalPriority.GpInfo {

	ceriteria := a.CreateCeriteria()
	alternative := a.CreateAlternatives()

	g := globalPriority.GlobalPriority(&ceriteria, &alternative, a.TableCandidate)

	return g
}
