package ahp

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (r Row) LenRow() int {
	return len(r)
}

func (r Row) LenCol() int {

	var lenCol int

	for _, val := range r {
		lenCol = len(val)
		break
	}

	return lenCol
}

func (r *Row) ProcessSum() Col {
	c := Col{}
	for _, val := range *r { //row
		for _, val2 := range val { //col
			for key3, val3 := range val2 {
				c[key3] += (val3)
				c[key3] = (c[key3])
			}
		}
	}
	return c
}

func (r *Row) ProcessEigen(a *Ahp) Row {
	res := Row{}
	for key, val := range *r { //row
		res[key] = []Col{}
		for _, val2 := range val { //col
			for key3, val3 := range val2 {
				res[key] = append(res[key], Col{
					key3: val3 / float64(a.Sum[key3]),
				})

			}
		}
	}

	return res
}

func (r *Row) ProcessAvg(a *Ahp) map[string]Info {

	type infoType struct {
		value float64
		count float64
		nums  []string
	}

	hash := make(map[string]infoType, r.LenCol())
	var info infoType

	for key, val := range r.ProcessEigen(a) {
		info = infoType{}
		for _, val2 := range val {
			for _, val3 := range val2 {
				info.count++
				info.value += val3
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

	return res
}

func (r *Row) ToJson() (string, error) {
	byte, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func (r *Ahp) ToJson() (string, error) {
	byte, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func NewTable() *Row {
	return &Row{}
}
