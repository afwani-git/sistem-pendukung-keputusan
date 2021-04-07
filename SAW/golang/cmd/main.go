package main

import (
	"net/http"

	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw"
	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw/normalized"
	"github.com/afwani-git/sistem-pendukung-keputusan/SAW/golang/pkg/saw/preferenceValCalc"
	"github.com/labstack/echo/v4"
)

func main() {
	criteriaTable := make(saw.CriteriaTable)

	const C1 = "Penguasaan Aspek teknis"
	criteriaTable[C1] = saw.CriteriaInfo{
		Name:  C1,
		Desc:  C1,
		Value: 1.8,
		Attr:  saw.BENEFIT,
	}

	const C2 = "Pengalaman Kerja"
	criteriaTable[C2] = saw.CriteriaInfo{
		Name:  C2,
		Desc:  C2,
		Value: 2.5,
		Attr:  saw.BENEFIT,
	}

	const C3 = "Interpersonal Skill"
	criteriaTable[C3] = saw.CriteriaInfo{
		Name:  C3,
		Desc:  C3,
		Value: 3.5,
		Attr:  saw.BENEFIT,
	}

	const C4 = "Usia"
	criteriaTable[C4] = saw.CriteriaInfo{
		Name:  C4,
		Desc:  C4,
		Value: 1.2,
		Attr:  saw.COST,
	}

	const C5 = "Status Perkawinan"
	criteriaTable[C5] = saw.CriteriaInfo{
		Name:  C5,
		Desc:  C5,
		Value: 2.8,
		Attr:  saw.COST,
	}

	alternativeTable := make(saw.AlternativeTable)

	const A1 saw.AlternativeName = "Wawan"
	alternativeTable[A1] = saw.AlternativeInfo{
		Name:          A1,
		CriteriaTable: make(map[saw.CriteriaName]saw.CriteriaAlternative),
	}
	//criteria
	alternativeTable[A1].CriteriaTable[C1] = saw.CriteriaAlternative{
		Values:   7,
		Criteria: criteriaTable[C1],
	}
	alternativeTable[A1].CriteriaTable[C2] = saw.CriteriaAlternative{
		Values:   4.5,
		Criteria: criteriaTable[C2],
	}
	alternativeTable[A1].CriteriaTable[C3] = saw.CriteriaAlternative{
		Values:   9,
		Criteria: criteriaTable[C3],
	}
	alternativeTable[A1].CriteriaTable[C4] = saw.CriteriaAlternative{
		Values:   37,
		Criteria: criteriaTable[C4],
	}
	alternativeTable[A1].CriteriaTable[C5] = saw.CriteriaAlternative{
		Values:   8,
		Criteria: criteriaTable[C5],
	}

	const A2 saw.AlternativeName = "Shinta"
	alternativeTable[A2] = saw.AlternativeInfo{
		Name:          A2,
		CriteriaTable: make(map[saw.CriteriaName]saw.CriteriaAlternative),
	}
	//criteria
	alternativeTable[A2].CriteriaTable[C1] = saw.CriteriaAlternative{
		Values:   7.5,
		Criteria: criteriaTable[C1],
	}
	alternativeTable[A2].CriteriaTable[C2] = saw.CriteriaAlternative{
		Values:   2,
		Criteria: criteriaTable[C2],
	}
	alternativeTable[A2].CriteriaTable[C3] = saw.CriteriaAlternative{
		Values:   6.5,
		Criteria: criteriaTable[C3],
	}
	alternativeTable[A2].CriteriaTable[C4] = saw.CriteriaAlternative{
		Values:   40,
		Criteria: criteriaTable[C4],
	}
	alternativeTable[A2].CriteriaTable[C5] = saw.CriteriaAlternative{
		Values:   8,
		Criteria: criteriaTable[C5],
	}

	const A3 saw.AlternativeName = "Gatot"
	alternativeTable[A3] = saw.AlternativeInfo{
		Name:          A3,
		CriteriaTable: make(map[saw.CriteriaName]saw.CriteriaAlternative),
	}
	//criteria
	alternativeTable[A3].CriteriaTable[C1] = saw.CriteriaAlternative{
		Values:   6,
		Criteria: criteriaTable[C1],
	}
	alternativeTable[A3].CriteriaTable[C2] = saw.CriteriaAlternative{
		Values:   5.5,
		Criteria: criteriaTable[C2],
	}
	alternativeTable[A3].CriteriaTable[C3] = saw.CriteriaAlternative{
		Values:   6,
		Criteria: criteriaTable[C3],
	}
	alternativeTable[A3].CriteriaTable[C4] = saw.CriteriaAlternative{
		Values:   34,
		Criteria: criteriaTable[C4],
	}
	alternativeTable[A3].CriteriaTable[C5] = saw.CriteriaAlternative{
		Values:   5,
		Criteria: criteriaTable[C5],
	}

	const A4 saw.AlternativeName = "Nina"
	alternativeTable[A4] = saw.AlternativeInfo{
		Name:          A4,
		CriteriaTable: make(map[saw.CriteriaName]saw.CriteriaAlternative),
	}
	//criteria
	alternativeTable[A4].CriteriaTable[C1] = saw.CriteriaAlternative{
		Values:   7.5,
		Criteria: criteriaTable[C1],
	}
	alternativeTable[A4].CriteriaTable[C2] = saw.CriteriaAlternative{
		Values:   0.5,
		Criteria: criteriaTable[C2],
	}
	alternativeTable[A4].CriteriaTable[C3] = saw.CriteriaAlternative{
		Values:   7.5,
		Criteria: criteriaTable[C3],
	}
	alternativeTable[A4].CriteriaTable[C4] = saw.CriteriaAlternative{
		Values:   23,
		Criteria: criteriaTable[C4],
	}
	alternativeTable[A4].CriteriaTable[C5] = saw.CriteriaAlternative{
		Values:   10,
		Criteria: criteriaTable[C5],
	}

	const A5 saw.AlternativeName = "Bella"
	alternativeTable[A5] = saw.AlternativeInfo{
		Name:          A5,
		CriteriaTable: make(map[saw.CriteriaName]saw.CriteriaAlternative),
	}
	//criteria
	alternativeTable[A5].CriteriaTable[C1] = saw.CriteriaAlternative{
		Values:   7,
		Criteria: criteriaTable[C1],
	}
	alternativeTable[A5].CriteriaTable[C2] = saw.CriteriaAlternative{
		Values:   2,
		Criteria: criteriaTable[C2],
	}
	alternativeTable[A5].CriteriaTable[C3] = saw.CriteriaAlternative{
		Values:   8.5,
		Criteria: criteriaTable[C3],
	}
	alternativeTable[A5].CriteriaTable[C4] = saw.CriteriaAlternative{
		Values:   44,
		Criteria: criteriaTable[C4],
	}
	alternativeTable[A5].CriteriaTable[C5] = saw.CriteriaAlternative{
		Values:   8,
		Criteria: criteriaTable[C5],
	}

	normalizedTable := normalized.NormalizedValue(criteriaTable, alternativeTable)
	calculatedTable := preferenceValCalc.CalcPreference(normalizedTable, criteriaTable)

	e := echo.New()

	e.GET("/criteria", func(c echo.Context) error {
		return c.JSON(http.StatusOK, criteriaTable)
	})

	e.GET("/alternative", func(c echo.Context) error {
		return c.JSON(http.StatusOK, alternativeTable)
	})

	e.GET("/normalized", func(c echo.Context) error {
		return c.JSON(http.StatusOK, normalizedTable)
	})

	e.GET("/calculated", func(c echo.Context) error {
		return c.JSON(http.StatusOK, calculatedTable)
	})

	e.Logger.Fatal(e.Start(":1323"))

}
