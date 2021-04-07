package saw

type CriteriaName string
type Attr string

const BENEFIT Attr = "benefit"
const COST Attr = "cost"

type CriteriaInfo struct {
	Name  CriteriaName `json:"name"`
	Desc  string       `json:"description"`
	Value float64      `json:"value"`
	Attr  Attr         `json:"attribute"`
}

type CriteriaTable map[CriteriaName]CriteriaInfo

type AlternativeName string

type CriteriaAlternative struct {
	Values   float64      `json:"values"`
	Criteria CriteriaInfo `json:"criteria"`
}

type AlternativeInfo struct {
	Name          AlternativeName                      `json:"name"`
	CriteriaTable map[CriteriaName]CriteriaAlternative `json:"criteria"`
}

type AlternativeTable map[AlternativeName]AlternativeInfo
