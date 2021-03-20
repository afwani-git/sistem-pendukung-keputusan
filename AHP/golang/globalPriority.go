package main

import (
	"sort"
	"strconv"
	"strings"
)

type Cgp struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

type InfoRank struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
	Key   string  `json:"key"`
}

type GpInfo struct {
	Table   map[string]Col
	Process map[string]Cgp   `json:"process"`
	Ranking map[int]InfoRank `json:"ranking"`
}

type Alternative struct {
	Desc   string `json:"description"`
	Key    string `json:"key"`
	Values *Ahp   `json:"list"`
}

func globalPriority(c *Ahp, a *[]Alternative) GpInfo {

	gpI := &GpInfo{}
	table := map[string]Col{}

	//extraction
	for _, val := range *a {

		c := Col{}

		for key2, val2 := range val.Values.Avg {
			c[key2] = val2.Value
		}

		table[val.Key] = c

	}

	type Info struct {
		val  float64
		nums []string
	}

	gpI.Table = table

	//helper
	aInfo := map[string][]Info{}

	for key, val := range gpI.Table {
		for key2, val2 := range val {
			// fmt.Println(key, key2)

			v1 := strconv.FormatFloat(val2, 'f', -3, 64)
			v2 := strconv.FormatFloat(c.Avg[key].Value, 'f', -3, 64)

			aInfo[key2] = append(aInfo[key2], Info{
				val:  val2 * c.Avg[key].Value,
				nums: []string{"( " + v1 + " x " + v2 + " )"},
			})

			// fmt.Println(aInfo)
		}
	}

	aInfo2 := map[string]Info{}
	for key, val := range aInfo {
		info := Info{}
		// *aInfo2[key] = Info{val: 0, nums: []string{}}
		for _, val2 := range val {
			// fmt.Println(key, key2, val2)
			info.val += val2.val
			info.nums = append(info.nums, val2.nums[0])
		}
		aInfo2[key] = info
	}

	process := map[string]Cgp{}
	for key, val := range aInfo2 {
		cgp := Cgp{
			Value: val.val,
			Desc:  strings.Join(val.nums, " + ") + " = " + strconv.FormatFloat(val.val, 'f', 3, 64),
		}
		process[key] = cgp
	}

	gpI.Process = process

	//sort / ranking

	values := make([]string, 0, len(gpI.Process))

	for _, val := range gpI.Process {

		vS := strconv.FormatFloat(val.Value, 'f', 3, 64)

		values = append(values, vS)
	}

	sort.StringSlice(values).Sort()
	ranking := make(map[int]InfoRank, len(gpI.Process))

	for key, val := range gpI.Process {
		sVal := strconv.FormatFloat(val.Value, 'f', 3, 64)
		no := sort.StringSlice(values).Search(sVal) + 1
		ranking[no] = InfoRank{
			Value: val.Value,
			Desc:  val.Desc,
			Key:   key,
		}
	}

	gpI.Ranking = ranking
	// fmt.Println(gpI)
	return *gpI
}
