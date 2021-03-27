package gap

type SubAspec string
type Aspec string
type CadidateName string

const CORE_FACTORY = "Core"
const SECONDARY_FACTORY = "Secondary"

type FactorType string

type SubAspectVal struct {
	Name        SubAspec   `json:"assessmentFactors"`
	TargetValue int        `json:"TargetValue"`
	Type        FactorType `json:"type"`
}

type AspectTargetVal struct {
	ApsecW      float64                   `json:"aspectWeight"`
	SecFactory  float64                   `json:"secodaryFactoryVal"`
	CoreFactory float64                   `json:"coreFactoryVal"`
	SubAspect   map[SubAspec]SubAspectVal `json:"subAspect"`
}

type AspectTable map[Aspec]AspectTargetVal

type CadidateVal struct {
	Value int      `json:"value"`
	Key   SubAspec `json:"key"`
}

type CadidateInfo struct {
	Name  CadidateName             `json:"name"`
	Value map[SubAspec]CadidateVal `json:"value"`
}

type CandidateTable map[Aspec]map[CadidateName]CadidateInfo

type InfoWeig struct {
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}

type WeigTable map[string]InfoWeig

type Gap struct {
	AspectTable   AspectTable    `json:"aspectTable"`
	CadidateTable CandidateTable `json:"candidateTable"`
	Wtable        WeigTable      `json:"weightTable"`
}
