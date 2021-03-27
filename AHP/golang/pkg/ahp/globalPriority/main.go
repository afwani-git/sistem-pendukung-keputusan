package globalPriority

import (
	"sort"
	"strconv"

	"github.com/afwani-git/sistem-pendukung-keputusan/AHP/golang/pkg/ahp"
)

func GlobalPriority(c *ahp.Ahp, a *[]Alternative) GpInfo {

	gpI := &GpInfo{}
	table := &map[string]ahp.Col{}

	//extraction
	extractValues(table, a)

	type Info struct {
		val  float64
		nums []string
	}

	gpI.Table = *table

	process := calculateValues(c, gpI)

	gpI.Process = process

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
	return *gpI
}
