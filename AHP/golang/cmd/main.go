package main

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/factory"
	"github.com/labstack/echo/v4"
)

func main() {

	//ceriteria table
	ceTable := ahp.NewTable()

	(*ceTable)["PERFORMA"] = []ahp.Col{{"PERFORMA": 1.0}, {"PENDIDIKAN": 3.0}, {"KOMUNIKASI": 5.0}}
	(*ceTable)["PENDIDIKAN"] = []ahp.Col{{"PERFORMA": 1.0 / 3.0}, {"PENDIDIKAN": 1.0}, {"KOMUNIKASI": 3.0}}
	(*ceTable)["KOMUNIKASI"] = []ahp.Col{{"PERFORMA": 1.0 / 5.0}, {"PENDIDIKAN": 1.0 / 3.0}, {"KOMUNIKASI": 1.0}}

	////////////////////////
	//alternative pendidikan
	////////////////////////
	//pendidikan alternative table
	pdTable := ahp.NewTable() //pendidikanTable

	(*pdTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 2.0}, {"C": 2.0}}
	(*pdTable)["B"] = []ahp.Col{{"A": 1.0 / 2.0}, {"B": 1}, {"C": 2.0}}
	(*pdTable)["C"] = []ahp.Col{{"A": 1.0 / 2.0}, {"B": 1 / 2.0}, {"C": 1.0}}

	pdAlternative := factory.AlternativeTable{
		Key:   "PENDIDIKAN",
		Desc:  "alternative pendidikan",
		Table: *pdTable,
	}

	//////////////////////
	//alternative performa
	//////////////////////
	//performa alternative table
	peTable := ahp.NewTable()

	(*peTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 3.0}, {"C": 4.0}}
	(*peTable)["B"] = []ahp.Col{{"A": 1.0 / 3.0}, {"B": 1}, {"C": 2.0}}
	(*peTable)["C"] = []ahp.Col{{"A": 1.0 / 4.0}, {"B": 1 / 2.0}, {"C": 1.0}}

	peAlternative := factory.AlternativeTable{
		Key:   "PERFORMA",
		Desc:  "alternative performa",
		Table: *peTable,
	}

	//////////////////////
	//alternative Komunikasi
	//////////////////////
	//motivasi alternative table
	maTable := ahp.NewTable()

	(*maTable)["A"] = []ahp.Col{{"A": 1.0}, {"B": 3.0}, {"C": 4.0}}
	(*maTable)["B"] = []ahp.Col{{"A": 1.0 / 3.0}, {"B": 1}, {"C": 2.0}}
	(*maTable)["C"] = []ahp.Col{{"A": 1.0 / 4.0}, {"B": 1 / 2.0}, {"C": 1.0}}

	maAlternative := factory.AlternativeTable{
		Key:   "KOMUNIKASI",
		Desc:  "alternative Komunikasi",
		Table: *maTable,
	}

	alternativesTable := []factory.AlternativeTable{
		pdAlternative, peAlternative, maAlternative,
	}

	//candidate table
	cTable := make(ahp.CandidateTable)
	//afwani
	cTable["Afwani"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Afwani"]["PENDIDIKAN"] = "A"
	cTable["Afwani"]["PERFORMA"] = "A"
	cTable["Afwani"]["KOMUNIKASI"] = "A"

	//dana
	cTable["Dana"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Dana"]["PENDIDIKAN"] = "B"
	cTable["Dana"]["PERFORMA"] = "B"
	cTable["Dana"]["KOMUNIKASI"] = "B"

	//jhon
	cTable["Jhon"] = make(map[ahp.AspecName]ahp.AspecValue)
	cTable["Jhon"]["PENDIDIKAN"] = "C"
	cTable["Jhon"]["PERFORMA"] = "A"
	cTable["Jhon"]["KOMUNIKASI"] = "A"

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
