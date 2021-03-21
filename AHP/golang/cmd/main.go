package main

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/consistency"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/globalPriority"
	"github.com/labstack/echo/v4"
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

	// alternative
	var alternatives []globalPriority.Alternative

	////////////////////////
	//alternative pendidikan
	////////////////////////
	pdAlternative := &ahp.Ahp{}
	//pendidikan alternative table
	pdTable := ahp.NewTable() //pendidikanTable

	(*pdTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 2.0}}
	(*pdTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*pdTable)["KC"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	pdAlternative.Table = *pdTable
	pdAlternative.Sum = pdTable.ProcessSum()
	pdAlternative.Eign = pdTable.ProcessEigen(pdAlternative)
	pdAlternative.Avg = pdTable.ProcessAvg(pdAlternative)

	//cek cosistency
	pdAlternative.Lmaks = consistency.ProcessLamdaMaks(pdAlternative)
	pdAlternative.Ci = consistency.Ci(pdAlternative.Lmaks, float64(3))

	ordo = consistency.Ordo(pdTable.LenCol())
	pdAlternative.Cr = consistency.Cr(pdAlternative.Ci, consistency.OrdoToRi(ordo))

	//append to alternative list
	alternatives = append(alternatives, globalPriority.Alternative{
		Desc:   "alternatif pendidikan",
		Key:    "PD",
		Values: pdAlternative,
	})

	//////////////////////
	//alternative performa
	//////////////////////
	peAlternative := &ahp.Ahp{}
	//performa alternative table
	peTable := ahp.NewTable()

	(*peTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*peTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 3.0}}
	(*peTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 3.0}, {"KC": 1.0}}

	peAlternative.Table = *peTable
	peAlternative.Sum = peTable.ProcessSum()
	peAlternative.Eign = peTable.ProcessEigen(peAlternative)
	peAlternative.Avg = peTable.ProcessAvg(peAlternative)

	//cek cosistency
	peAlternative.Lmaks = consistency.ProcessLamdaMaks(peAlternative)
	peAlternative.Ci = consistency.Ci(peAlternative.Lmaks, float64(peAlternative.Table.LenCol()))

	ordo = consistency.Ordo(pdTable.LenCol())
	peAlternative.Cr = consistency.Cr(peAlternative.Ci, consistency.OrdoToRi(ordo))

	//append to alternative list
	alternatives = append(alternatives, globalPriority.Alternative{
		Desc:   "alternatif performa",
		Key:    "PE",
		Values: peAlternative,
	})

	//////////////////////
	//alternative motivasi
	//////////////////////
	maAlternative := &ahp.Ahp{}
	//motivasi alternative table
	maTable := ahp.NewTable()

	(*maTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*maTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*maTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	maAlternative.Table = *maTable
	maAlternative.Sum = maTable.ProcessSum()
	maAlternative.Eign = maTable.ProcessEigen(maAlternative)
	maAlternative.Avg = maTable.ProcessAvg(maAlternative)

	//cek cosistency
	maAlternative.Lmaks = consistency.ProcessLamdaMaks(maAlternative)
	maAlternative.Ci = consistency.Ci(maAlternative.Lmaks, float64(maAlternative.Table.LenCol()))

	ordo = consistency.Ordo(pdTable.LenCol())
	maAlternative.Cr = consistency.Cr(maAlternative.Ci, consistency.OrdoToRi(ordo))

	//append to alternative list
	alternatives = append(alternatives, globalPriority.Alternative{
		Desc:   "alternatif Motivasi",
		Key:    "MA",
		Values: maAlternative,
	})

	//////////////////////
	//alternative komunikasi
	//////////////////////
	kaAlternative := &ahp.Ahp{}
	//komunikasi alternative table
	kaTable := ahp.NewTable()

	(*kaTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*kaTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*kaTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	kaAlternative.Table = *kaTable
	kaAlternative.Sum = kaTable.ProcessSum()
	kaAlternative.Eign = kaTable.ProcessEigen(kaAlternative)
	kaAlternative.Avg = kaTable.ProcessAvg(kaAlternative)

	//cek cosistency
	kaAlternative.Lmaks = consistency.ProcessLamdaMaks(kaAlternative)
	kaAlternative.Ci = consistency.Ci(kaAlternative.Lmaks, float64(kaAlternative.Table.LenCol()))

	ordo = consistency.Ordo(pdTable.LenCol())
	kaAlternative.Cr = consistency.Cr(kaAlternative.Ci, consistency.OrdoToRi(ordo))

	//append to alternative list
	alternatives = append(alternatives, globalPriority.Alternative{
		Desc:   "alternatif Motivasi",
		Key:    "MA",
		Values: kaAlternative,
	})

	g := globalPriority.GlobalPriority(ceriteria, &alternatives)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, alternatives)
	})

	e.GET("/ranking", func(c echo.Context) error {
		return c.JSON(200, g)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
