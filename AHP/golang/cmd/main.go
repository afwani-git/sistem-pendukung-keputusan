package main

import (
	"fmt"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp/factory"
	"github.com/labstack/echo/v4"
)

func main() {

	//ceriteria table
	ceTable := ahp.NewTable()

	(*ceTable)["PD"] = []ahp.Col{{"PD": 1.0}, {"PF": 3.0}, {"KB": 3.0}, {"MA": 5}}
	(*ceTable)["PF"] = []ahp.Col{{"PD": 1.0 / 3.0}, {"PF": 1.0}, {"KB": 3.0}, {"MA": 3.0}}
	(*ceTable)["KB"] = []ahp.Col{{"PD": 1.0 / 3.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0}, {"MA": 3.0}}
	(*ceTable)["MA"] = []ahp.Col{{"PD": 1.0 / 5.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0 / 3}, {"MA": 1.0}}

	////////////////////////
	//alternative pendidikan
	////////////////////////
	// pdAlternative := &ahp.Ahp{}
	//pendidikan alternative table
	pdTable := ahp.NewTable() //pendidikanTable

	(*pdTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 2.0}}
	(*pdTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*pdTable)["KC"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	pdAlternative := factory.AlternativeTable{
		Key:   "PD",
		Desc:  "alternative pendidikan",
		Table: *pdTable,
	}

	//////////////////////
	//alternative performa
	//////////////////////
	// peAlternative := &ahp.Ahp{}
	//performa alternative table
	peTable := ahp.NewTable()

	(*peTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*peTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 3.0}}
	(*peTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 3.0}, {"KC": 1.0}}

	peAlternative := factory.AlternativeTable{
		Key:   "PE",
		Desc:  "alternative performa",
		Table: *peTable,
	}

	//////////////////////
	//alternative motivasi
	//////////////////////
	// maAlternative := &ahp.Ahp{}
	//motivasi alternative table
	maTable := ahp.NewTable()

	(*maTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*maTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*maTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	maAlternative := factory.AlternativeTable{
		Key:   "MA",
		Desc:  "alternative motivasi",
		Table: *maTable,
	}

	//////////////////////
	//alternative komunikasi
	//////////////////////
	// kaAlternative := &ahp.Ahp{}
	//komunikasi alternative table
	kaTable := ahp.NewTable()

	(*kaTable)["KA"] = []ahp.Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	(*kaTable)["KB"] = []ahp.Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	(*kaTable)["KC"] = []ahp.Col{{"KA": 1.0 / 3.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	kaAlternative := factory.AlternativeTable{
		Key:   "KA",
		Desc:  "alternative KA",
		Table: *kaTable,
	}

	alternativesTable := []factory.AlternativeTable{
		pdAlternative, peAlternative, kaAlternative, maAlternative,
	}

	//factory
	AhpFactory := factory.AhpFactory{
		TableCriteria:    ceTable,
		TableAlternative: &alternativesTable,
	}

	fmt.Println(AhpFactory.CreateAlternatives())

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, AhpFactory.CreateAlternatives())
	})

	e.GET("/ranking", func(c echo.Context) error {
		return c.JSON(200, AhpFactory.Create())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
