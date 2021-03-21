package ahp

type LambdaInfo struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

type LamdaMaks struct {
	Process map[string]LambdaInfo `json:"process"`
	Total   LambdaInfo            `json:"total"`
}

type Info struct {
	Value float64 `json:"value"`
	Desc  string  `json:"desc"`
}

type CiInfo struct {
	Value float64 `json:"value"`
	Desc  string  `json:"description"`
}

type CrInfo struct {
	Value  float64 `json:"value"`
	Desc   string  `json:"description"`
	Consis bool    `json:"consistance"`
}

type Col map[string]float64
type Row map[string][]Col
type Wow string

type Ahp struct {
	Table Row             `json:"table"`
	Sum   Col             `json:"total"`
	Eign  Row             `json:"eigen"`
	Avg   map[string]Info `json:"avg"`
	Lmaks LamdaMaks       `json:"lambdaMaks"`
	Ci    CiInfo          `json:"ci"`
	Cr    CrInfo          `json:"cr"`
}

func (w *Wow) Tes() {

}
