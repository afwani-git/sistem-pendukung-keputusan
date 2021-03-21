package main

import (
	"fmt"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/consistency"
)

func main() {

	//ceriteria
	ceriteria := &ahp.Ahp{}

	//ceriteria table
	ceTable := ahp.NewTable()

	(*ceTable)["PD"] = []ahp.Col{{"PD": 1.0}, {"PF": 3.0}, {"KB": 3.0}, {"MA": 5}}
	(*ceTable)["PF"] = []ahp.Col{{"PD": 1.0 / 3.0}, {"PF": 1.0}, {"KB": 3.0}, {"MA": 3.0}}
	(*ceTable)["KB"] = []ahp.Col{{"PD": 1.0 / 3.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0}, {"MA": 3.0}}
	(*ceTable)["MA"] = []ahp.Col{{"PD": 1.0 / 5.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0 / 3}, {"MA": 1.0}}

	ceriteria.Table = *ceTable
	ceriteria.Sum = ceTable.ProcessSum()
	ceriteria.Eign = ceTable.ProcessEigen(ceriteria)
	ceriteria.Avg = ceTable.ProcessAvg(ceriteria)

	//cek cosistency
	ceriteria.Lmaks = consistency.ProcessLamdaMaks(ceriteria)
	ceriteria.Ci = consistency.Ci(ceriteria.Lmaks, float64(ceriteria.Table.LenCol()))

	ordo := consistency.Ordo(ceTable.LenCol())
	ceriteria.Cr = consistency.Cr(ceriteria.Ci, consistency.OrdoToRi(ordo))
	//process

	json, err := ceriteria.ToJson()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json)
	// fmt.Println(ceriteria)
}
