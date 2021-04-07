package main

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/factory"
	"github.com/labstack/echo/v4"
)

func main() {

	//ceriteria table
	ceTable := ahp.NewTable()

	(*ceTable)["ABSENSI"] = []ahp.Col{{"ABSENSI": 1.0}, {"IPK": 2.0}, {"PRILAKU": 4.0}}
	(*ceTable)["IPK"] = []ahp.Col{{"ABSENSI": 1.0 / 2.0}, {"IPK": 1.0}, {"PRILAKU": 4.0}}
	(*ceTable)["PRILAKU"] = []ahp.Col{{"ABSENSI": 1.0 / 4.0}, {"IPK": 1.0 / 4.0}, {"PRILAKU": 1.0}}

	////////////////////////
	//alternative IPK
	////////////////////////
	//IPK alternative table
	pdTable := ahp.NewTable() //IPKTable

	(*pdTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 3.0}, {"C": 4.0}}
	(*pdTable)["B"] = []ahp.Col{{"A": 1.0 / 3.0}, {"B": 1}, {"C": 3.0}}
	(*pdTable)["C"] = []ahp.Col{{"A": 1.0 / 4.0}, {"B": 1 / 3.0}, {"C": 1.0}}

	pdAlternative := factory.AlternativeTable{
		Key:   "IPK",
		Desc:  "alternative IPK",
		Table: *pdTable,
	}

	//////////////////////
	//alternative ABSENSI
	//////////////////////
	//ABSENSI alternative table
	peTable := ahp.NewTable()

	(*peTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 5.0}, {"C": 7.0}}
	(*peTable)["B"] = []ahp.Col{{"A": 1.0 / 5.0}, {"B": 1}, {"C": 3.0}}
	(*peTable)["C"] = []ahp.Col{{"A": 1.0 / 7.0}, {"B": 1 / 3.0}, {"C": 1.0}}

	peAlternative := factory.AlternativeTable{
		Key:   "ABSENSI",
		Desc:  "alternative ABSENSI",
		Table: *peTable,
	}

	//////////////////////
	//alternative PRILAKU
	//////////////////////
	//motivasi alternative table
	maTable := ahp.NewTable()

	(*maTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 2.0}, {"C": 2.0}}
	(*maTable)["B"] = []ahp.Col{{"A": 1.0 / 2.0}, {"B": 1}, {"C": 2.0}}
	(*maTable)["C"] = []ahp.Col{{"A": 1.0 / 2.0}, {"B": 1 / 2.0}, {"C": 1.0}}

	maAlternative := factory.AlternativeTable{
		Key:   "PRILAKU",
		Desc:  "alternative PRILAKU",
		Table: *maTable,
	}

	alternativesTable := []factory.AlternativeTable{
		pdAlternative, peAlternative, maAlternative,
	}

	//candidate table
	cTable := make(ahp.CandidateTable)
	//afwani
	cTable["Afwani"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Afwani"]["IPK"] = "A"
	cTable["Afwani"]["ABSENSI"] = "A"
	cTable["Afwani"]["PRILAKU"] = "A"

	//dana
	cTable["Dana"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Dana"]["IPK"] = "B"
	cTable["Dana"]["ABSENSI"] = "B"
	cTable["Dana"]["PRILAKU"] = "B"

	//jhon
	cTable["Jhon"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Jhon"]["IPK"] = "C"
	cTable["Jhon"]["ABSENSI"] = "A"
	cTable["Jhon"]["PRILAKU"] = "A"

	//factory
	AhpFactory := factory.AhpFactory{
		TableCriteria:    ceTable,
		TableAlternative: &alternativesTable,
		TableCandidate:   cTable,
	}

	e := echo.New()

	e.GET("/candidate", func(c echo.Context) error {
		return c.JSON(200, cTable)
	})

	e.GET("/creteria", func(c echo.Context) error {
		return c.JSON(200, AhpFactory.CreateCeriteria())
	})

	e.GET("/alternative", func(c echo.Context) error {
		return c.JSON(200, AhpFactory.CreateAlternatives())
	})

	e.GET("/ranking", func(c echo.Context) error {
		return c.JSON(200, AhpFactory.Create())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
