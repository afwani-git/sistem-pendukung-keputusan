package main

import (
	"strconv"
	"strings"
)

type LambdaInfo struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

type LamdaMaks struct {
	Process map[string]LambdaInfo `json:"process"`
	Total   LambdaInfo            `json:"total"`
}

func processLamdaMaks(a *Ahp) LamdaMaks {

	//make hash/key-val table
	hash := make(map[string]LambdaInfo, len(a.Sum))
	// key := make([]string, len(a.Sum))

	for key, val := range a.Sum {

		valStr1 := strconv.FormatFloat(val, 'f', 3, 64)
		valStr2 := strconv.FormatFloat(a.Avg[key].Value, 'f', 3, 64)

		hash[key] = LambdaInfo{
			Value: ((val) * (a.Avg[key].Value)),
			Desc:  "( " + valStr1 + " * " + valStr2 + " )",
		}
	}

	tmpDesc := make([]string, 0, len(a.Sum))
	total := LambdaInfo{}
	for _, val := range hash {
		tmpDesc = append(tmpDesc, val.Desc)
		total.Value += (val.Value)
		total.Value = (total.Value)
	}

	total.Desc = strings.Join(tmpDesc, " + ") + " = " + strconv.FormatFloat(total.Value, 'f', 3, 64)

	result := LamdaMaks{
		Process: hash,
		Total:   total,
	}

	return result
}
