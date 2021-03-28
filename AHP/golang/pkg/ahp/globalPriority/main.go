package globalPriority

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func GlobalPriority(c *ahp.Ahp, a *[]Alternative, cA ahp.CandidateTable) GpInfo {

	gpI := &GpInfo{}

	table := joinAllEign(a)
	gpI.Table = table

	covTable := convert(*c, gpI.Table, cA)

	gpI.ConvertedTable = covTable

	return *gpI
}
