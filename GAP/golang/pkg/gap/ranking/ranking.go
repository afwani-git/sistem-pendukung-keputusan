package ranking

import (
	"strconv"
	"strings"

	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/calculate"
)

func Ranking(t calculate.TableTotalCalcFactory, g gap.SpecTable) RankingTable {

	type calCulateAspec map[gap.Aspec]map[gap.CadidateName]struct {
		Value float64
		Desc  string
	}

	table := RankingTable{}
	clcAspec := make(calCulateAspec)

	for aspec, val := range t {
		clcAspec[aspec] = make(map[gap.CadidateName]struct {
			Value float64
			Desc  string
		})
		for cName, val2 := range val {

			clcAspec[aspec][cName] = struct {
				Value float64
				Desc  string
			}{
				Value: (g.AspectTable[aspec].ApsecW / 100) * val2.TotalValue.Value,
				Desc:  " ( " + strconv.FormatFloat((g.AspectTable[aspec].ApsecW/100), 'f', -1, 64) + " * " + strconv.FormatFloat((val2.TotalValue.Value), 'f', -1, 64) + " ) ",
			}
		}
	}

	table.Process = make(ProcessInfo)
	acumulate := make(map[gap.CadidateName][]string)

	for _, val := range clcAspec {
		for cName, val2 := range val {

			acumulate[cName] = append(acumulate[cName], val2.Desc)

			table.Process[cName] = ProcessValue{
				Name:        string(cName),
				FinalScore:  table.Process[cName].FinalScore + val2.Value,
				Description: strings.Join(acumulate[cName], " + "),
			}
		}
	}

	return table
}
