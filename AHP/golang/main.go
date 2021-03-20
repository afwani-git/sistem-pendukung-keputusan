package main

import "github.com/labstack/echo/v4"

//TODO: mari belajar package brader

func main() {

	ce := &Ahp{}

	// init
	ce.Table = Row{}

	ce.Table["PD"] = []Col{{"PD": 1.0}, {"PF": 3.0}, {"KB": 3.0}, {"MA": 5}}
	ce.Table["PF"] = []Col{{"PD": 1.0 / 3.0}, {"PF": 1.0}, {"KB": 3.0}, {"MA": 3.0}}
	ce.Table["KB"] = []Col{{"PD": 1.0 / 3.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0}, {"MA": 3.0}}
	ce.Table["MA"] = []Col{{"PD": 1.0 / 5.0}, {"PF": 1.0 / 3.0}, {"KB": 1.0 / 3}, {"MA": 1.0}}

	ce.processSum()
	ce.processEigen()
	ce.processAvg()

	lmaksCe := processLamdaMaks(ce)

	//validasi
	ciCe := Ci(lmaksCe, float64(len(ce.Sum)))
	crCe := Cr(ciCe, Ri(0.90))

	ce.Ci = ciCe
	ce.Cr = crCe
	ce.Lmaks = lmaksCe
	////////////////////////////
	// Alternative
	///////////////////////////

	alternatives := make([]Alternative, 0, len(ce.Sum))

	/////////////////////////////////////////////////
	//alternatif pendidikan
	////////////////////////////////////////////////
	p := &Ahp{}

	// init
	p.Table = Row{}

	p.Table["KA"] = []Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 2.0}}
	p.Table["KB"] = []Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	p.Table["KC"] = []Col{{"KA": 1.0 / 2.0}, {"KB": 1 / 2.0}, {"KC": 1.0}}

	p.processSum()
	p.processEigen()
	p.processAvg()

	lmaksP := processLamdaMaks(p)

	//validasi
	ciP := Ci(lmaksP, float64(len(p.Sum)))
	crP := Cr(ciP, Ri(0.58))

	p.Ci = ciP
	p.Cr = crP
	p.Lmaks = lmaksP

	alternatives = append(alternatives, Alternative{
		Desc:   "alternatif pendidikan",
		Key:    "PD",
		Values: p,
	})

	/////////////////////////////////////////////////
	//alternatif performa
	////////////////////////////////////////////////
	pe := &Ahp{}

	// init
	pe.Table = Row{}

	//isi table
	pe.Table["KA"] = []Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	pe.Table["KB"] = []Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 3.0}}
	pe.Table["KC"] = []Col{{"KA": 1.0 / 2.0}, {"KB": 1.0 / 3.0}, {"KC": 1.0}}

	pe.processSum()
	pe.processEigen()
	pe.processAvg()

	lmaksPe := processLamdaMaks(pe)

	//validasi
	ciPe := Ci(lmaksPe, float64(len(pe.Sum)))
	crPe := Cr(ciPe, Ri(0.58))

	pe.Ci = ciPe
	pe.Cr = crPe
	pe.Lmaks = lmaksPe

	alternatives = append(alternatives, Alternative{
		Desc:   "alternatif performa",
		Key:    "PF",
		Values: pe,
	})

	/////////////////////////////////////////////////
	//alternatif motivasi
	////////////////////////////////////////////////
	ma := &Ahp{}

	// init
	ma.Table = Row{}

	//isi table
	ma.Table["KA"] = []Col{{"KA": 1.0}, {"KB": 2.0}, {"KC": 3.0}}
	ma.Table["KB"] = []Col{{"KA": 1.0 / 2.0}, {"KB": 1}, {"KC": 2.0}}
	ma.Table["KC"] = []Col{{"KA": 1.0 / 3.0}, {"KB": 1.0 / 2.0}, {"KC": 1.0}}

	ma.processSum()
	ma.processEigen()
	ma.processAvg()

	lmaksMa := processLamdaMaks(ma)

	//validasi
	ciMa := Ci(lmaksMa, float64(len(ma.Sum)))
	crMa := Cr(ciMa, Ri(0.58))

	ma.Ci = ciMa
	ma.Cr = crMa
	ma.Lmaks = lmaksMa

	alternatives = append(alternatives, Alternative{
		Desc:   "alternatif Motivasi",
		Key:    "MA",
		Values: ma,
	})

	/////////////////////////////////////////////////
	//alternatif komunikasi
	////////////////////////////////////////////////
	ka := &Ahp{}

	// init
	ka.Table = Row{}

	//isi table
	ka.Table["KA"] = []Col{{"KA": 1.0}, {"KB": 3.0}, {"KC": 3.0}}
	ka.Table["KB"] = []Col{{"KA": 1.0 / 3.0}, {"KB": 1.0}, {"KC": 2.0}}
	ka.Table["KC"] = []Col{{"KA": 1.0 / 3.0}, {"KB": 1.0 / 2.0}, {"KC": 1.0}}

	ka.processSum()
	ka.processEigen()
	ka.processAvg()

	lmaksKa := processLamdaMaks(ka)

	//validasi
	ciKa := Ci(lmaksKa, float64(len(ka.Sum)))
	crKa := Cr(ciKa, Ri(0.58))

	ka.Ci = ciKa
	ka.Cr = crKa
	ka.Lmaks = lmaksKa

	alternatives = append(alternatives, Alternative{
		Desc:   "alternatif Komunikasi",
		Key:    "KB",
		Values: ka,
	})
	/////////////////////////////////////////////////
	//global priotity
	////////////////////////////////////////////////
	g := globalPriority(ce, &alternatives)

	//echo web

	//TODO: mari pisah ini broder
	e := echo.New()

	e.GET("/keriteria", func(c echo.Context) error {
		return c.JSON(200, ce)
	})

	e.GET("/alternatives", func(c echo.Context) error {
		return c.JSON(200, alternatives)
	})

	e.GET("/ranking", func(c echo.Context) error {
		return c.JSON(200, g)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
