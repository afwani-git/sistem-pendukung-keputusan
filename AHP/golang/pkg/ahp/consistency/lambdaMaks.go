package consistency

import (
	"strconv"
	"strings"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func ProcessLamdaMaks(a *ahp.Ahp) ahp.LamdaMaks {

	//make hash/key-val table
	hash := make(map[string]ahp.LambdaInfo, len(a.Sum))

	for key, val := range a.Sum {

		valStr1 := strconv.FormatFloat(val, 'f', 3, 64)
		valStr2 := strconv.FormatFloat(a.Avg[key].Value, 'f', 3, 64)

		hash[key] = ahp.LambdaInfo{
			Value: ((val) * (a.Avg[key].Value)),
			Desc:  "( " + valStr1 + " * " + valStr2 + " )",
		}
	}

	tmpDesc := make([]string, 0, len(a.Sum))
	total := ahp.LambdaInfo{}
	for _, val := range hash {
		tmpDesc = append(tmpDesc, val.Desc)
		total.Value += (val.Value)
		total.Value = (total.Value)
	}

	total.Desc = strings.Join(tmpDesc, " + ") + " = " + strconv.FormatFloat(total.Value, 'f', 3, 64)

	result := ahp.LamdaMaks{
		Process: hash,
		Total:   total,
	}

	return result
}
