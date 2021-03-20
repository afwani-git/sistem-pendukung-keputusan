package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Col map[string]float64

type Row map[string][]Col

type Info struct {
	Value float64 `json:"value"`
	Desc  string  `json:"desc"`
}

type Ahp struct {
	Table Row             `json:"table"`
	Sum   Col             `json:"total"`
	Eign  Row             `json:"eigen"`
	Avg   map[string]Info `json:"avg"`
	Lmaks LamdaMaks       `json:"lambdaMaks"`
	Ci    CiInfo          `json:"ci"`
	Cr    CrInfo          `json:"cr"`
}

func (t *Ahp) toJson() (string, error) {
	byte, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

//FIXME: fixing print()
func (t *Row) print() {
	drawTable := tablewriter.NewWriter(os.Stdout)

	var (
		h []string   //header
		r [][]string //rows
	)

	//set  header
	h = append(h, "KERITERIA")
	for row, col := range *t {
		h = append(h, row)
		newCol := []string{}
		newCol = append(newCol, row)
		for _, ir := range col {
			for _, v := range ir {
				newCol = append(newCol, strconv.FormatFloat(v, 'f', 3, 64))
			}
		}
		r = append(r, newCol)
	}

	drawTable.SetHeader(h)
	drawTable.AppendBulk(r)
	drawTable.Render()
}

func (t *Ahp) processSum() {
	c := Col{}
	for _, val := range t.Table { //row
		for _, val2 := range val { //col
			for key3, val3 := range val2 {
				c[key3] += (val3)
				c[key3] = (c[key3])
			}
		}
	}

	t.Sum = c
}

func (t *Ahp) processEigen() { // O(n^3)

	r := Row{}
	for key, val := range t.Table { //row
		r[key] = []Col{}
		for _, val2 := range val { //col
			for key3, val3 := range val2 {
				r[key] = append(r[key], Col{
					key3: ((val3) / (t.Sum[key3])),
				})

			}
		}
	}

	t.Eign = r
}

func (t *Ahp) processAvg() {
	type info struct {
		value float64
		count float64
		nums  []string
	}

	hash := make(map[string]info, len(t.Sum))

	for key, val := range t.Eign {
		info := info{}
		for _, val2 := range val {
			for _, val3 := range val2 {
				info.count++
				info.value += (val3)
				info.value = (info.value)
				info.nums = append(info.nums, strconv.FormatFloat(val3, 'f', 3, 64))
				hash[key] = info
			}
		}
	}

	res := make(map[string]Info)

	for key, val := range hash {
		res[key] = Info{
			Value: ((val.value) / (val.count)),
			Desc:  "( " + strings.Join(val.nums, " + ") + " ) / " + strconv.FormatFloat(val.count, 'f', -1, 64),
		}
	}

	t.Avg = res
}
