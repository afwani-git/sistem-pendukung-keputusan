package factory

import (
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/calculate"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/gapMapping"
	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/ranking"
)

func CreacteGap(s gap.SpecTable) ResultGap {

	result := ResultGap{}
	result.GapSpec = s
	result.GapMapping = gapMapping.MappingGap(s)
	result.GapCalculate = calculate.CalculateFactor(result.GapMapping, s)
	result.GapRanking = ranking.Ranking(result.GapCalculate, s)

	return result
}

func LoadDummy() ResultGap {

	//init
	aspectTable := make(gap.AspectTable)

	//KECERDASAN/////////////////////////////////////////////
	const KECERDASAN = "Kecerdasan"
	aspectTable[KECERDASAN] = gap.AspectTargetVal{
		ApsecW:      30,
		SecFactory:  30,
		CoreFactory: 70,
		SubAspect:   make(map[gap.SubAspec]gap.SubAspectVal),
	}

	//Kecerdasan => Common sense
	const COMMON_SENSE = "Common Sense"
	// aspectTable[KECERDASAN].SubAspect := make(map[gap.SubAspec]gap.SubAspectVal)
	aspectTable[KECERDASAN].SubAspect[COMMON_SENSE] = gap.SubAspectVal{
		Name:        COMMON_SENSE,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}
	//Kecerdasan => Verbalisasi Ide
	const VERBALISASI_IDE = "Verbalisasi Ide"
	aspectTable[KECERDASAN].SubAspect[VERBALISASI_IDE] = gap.SubAspectVal{
		Name:        VERBALISASI_IDE,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}
	//Kecerdasan => Sistematika Berpikir
	const SISTEMATIKA_BERPIKIR = "Sistematika Berpikir"
	aspectTable[KECERDASAN].SubAspect[SISTEMATIKA_BERPIKIR] = gap.SubAspectVal{
		Name:        SISTEMATIKA_BERPIKIR,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}
	//Kecerdasan => Penalaran dan Solusi Real
	const PENALARAN_DAN_SOLUSI_REAL = "Penalaran dan Solusi Real"
	aspectTable[KECERDASAN].SubAspect[PENALARAN_DAN_SOLUSI_REAL] = gap.SubAspectVal{
		Name:        PENALARAN_DAN_SOLUSI_REAL,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}
	//Kecerdasan => Konsentrasi
	const KONSENTRASI = "Konsentrasi"
	aspectTable[KECERDASAN].SubAspect[KONSENTRASI] = gap.SubAspectVal{
		Name:        KONSENTRASI,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}
	//Kecerdasan => Logika Praktis
	const LOGIKA_PRAKTIS = "Logika Praktis"
	aspectTable[KECERDASAN].SubAspect[LOGIKA_PRAKTIS] = gap.SubAspectVal{
		Name:        LOGIKA_PRAKTIS,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}
	//Kecerdasan => Fleksibilitas berpikir
	const FLEKSIBILITAS_BERPIKIR = "Fleksibilitas berpikir"
	aspectTable[KECERDASAN].SubAspect[FLEKSIBILITAS_BERPIKIR] = gap.SubAspectVal{
		Name:        FLEKSIBILITAS_BERPIKIR,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}
	//Kecerdasan => Imajinasi kreatif
	const IMAJINASI_KREATIF = "Imajinasi kreatif"
	aspectTable[KECERDASAN].SubAspect[IMAJINASI_KREATIF] = gap.SubAspectVal{
		Name:        IMAJINASI_KREATIF,
		TargetValue: 5,
		Type:        gap.CORE_FACTORY,
	}
	//Kecerdasan => Antisipasi
	const ANTISIPASI = "Antisipasi"
	aspectTable[KECERDASAN].SubAspect[ANTISIPASI] = gap.SubAspectVal{
		Name:        ANTISIPASI,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}
	//Kecerdasan => Potensi Kecerdasan
	const POTENSI_KECERDASAN = "Potensi Kecerdasan"
	aspectTable[KECERDASAN].SubAspect[POTENSI_KECERDASAN] = gap.SubAspectVal{
		Name:        POTENSI_KECERDASAN,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}
	//////////////////////////////////////////////////////////////////
	//sikap kerja////////////////////////////////////////////////////
	const SIKAP_KERJA = "Sikap kerja"
	aspectTable[SIKAP_KERJA] = gap.AspectTargetVal{
		ApsecW:      30,
		SecFactory:  35,
		CoreFactory: 65,
		SubAspect:   make(map[gap.SubAspec]gap.SubAspectVal),
	}
	//sikap kerja => energi psikis
	const ENERGI_PSIKIS = "Energi psikis"
	aspectTable[SIKAP_KERJA].SubAspect[ENERGI_PSIKIS] = gap.SubAspectVal{
		Name:        ENERGI_PSIKIS,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}

	//sikap kerja => Ketelitian dan tanggung jawab
	const KETELITIAN = "Ketelitian dan tanggung jawab"
	aspectTable[SIKAP_KERJA].SubAspect[KETELITIAN] = gap.SubAspectVal{
		Name:        KETELITIAN,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}

	//sikap kerja => Kehati-hatian
	const KEHATI_HATIAN = "Kehati-hatian"
	aspectTable[SIKAP_KERJA].SubAspect[KEHATI_HATIAN] = gap.SubAspectVal{
		Name:        KEHATI_HATIAN,
		TargetValue: 2,
		Type:        gap.SECONDARY_FACTORY,
	}

	//sikap kerja => Pengendalian Perasaan
	const PENGENDALIAN_PERASAAN = "Pengendalian Perasaan"
	aspectTable[SIKAP_KERJA].SubAspect[PENGENDALIAN_PERASAAN] = gap.SubAspectVal{
		Name:        PENGENDALIAN_PERASAAN,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}

	//sikap kerja => Dorongan Berprestasi
	const DORONGAN_PRESTASI = "Dorongan Berprestasi"
	aspectTable[SIKAP_KERJA].SubAspect[DORONGAN_PRESTASI] = gap.SubAspectVal{
		Name:        DORONGAN_PRESTASI,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}

	//sikap kerja => Vitalitas dan Perencanaan
	const VITALITAS_PERENCANAAN = "Vitalitas dan Perencanaan"
	aspectTable[SIKAP_KERJA].SubAspect[VITALITAS_PERENCANAAN] = gap.SubAspectVal{
		Name:        VITALITAS_PERENCANAAN,
		TargetValue: 5,
		Type:        gap.CORE_FACTORY,
	}
	//////////////////////////////////////////////////////////////////
	//Perilaku///////////////////////////////////////////////////////
	const PERILAKU = "Perilaku"
	aspectTable[PERILAKU] = gap.AspectTargetVal{
		ApsecW:      40,
		SecFactory:  30,
		CoreFactory: 70,
		SubAspect:   make(map[gap.SubAspec]gap.SubAspectVal),
	}
	//Perilaku => Dominance
	const DOMINACE = "Dominace"
	aspectTable[PERILAKU].SubAspect[DOMINACE] = gap.SubAspectVal{
		TargetValue: 3,
		Name:        DOMINACE,
		Type:        gap.SECONDARY_FACTORY,
	}

	//Perilaku => Influences
	const INFLUENCES = "Influences"
	aspectTable[PERILAKU].SubAspect[INFLUENCES] = gap.SubAspectVal{
		Name:        INFLUENCES,
		TargetValue: 3,
		Type:        gap.SECONDARY_FACTORY,
	}

	//Perilaku => Steadiness
	const STEADINESS = "Steadiness"
	aspectTable[PERILAKU].SubAspect[STEADINESS] = gap.SubAspectVal{
		Name:        STEADINESS,
		TargetValue: 4,
		Type:        gap.CORE_FACTORY,
	}

	//Perilaku => Steadiness
	const COMPLIANCE = "Compliance"
	aspectTable[PERILAKU].SubAspect[COMPLIANCE] = gap.SubAspectVal{
		Name:        COMPLIANCE,
		TargetValue: 5,
		Type:        gap.CORE_FACTORY,
	}

	candidateTable := make(gap.CandidateTable)
	//kecerdasan
	candidateTable[KECERDASAN] = make(map[gap.CadidateName]gap.CadidateInfo)
	candidateTable[SIKAP_KERJA] = make(map[gap.CadidateName]gap.CadidateInfo)
	candidateTable[PERILAKU] = make(map[gap.CadidateName]gap.CadidateInfo)
	//Usman
	const USMAN = "Usman"
	candidateTable[KECERDASAN][USMAN] = gap.CadidateInfo{
		Name: USMAN,
		Value: map[gap.SubAspec]gap.CadidateVal{
			COMMON_SENSE: {
				Value: 3,
				Key:   COMMON_SENSE,
			},
			VERBALISASI_IDE: {
				Value: 4,
				Key:   VERBALISASI_IDE,
			},
			SISTEMATIKA_BERPIKIR: {
				Value: 3,
				Key:   SISTEMATIKA_BERPIKIR,
			},
			PENALARAN_DAN_SOLUSI_REAL: {
				Value: 3,
				Key:   PENALARAN_DAN_SOLUSI_REAL,
			},
			KONSENTRASI: {
				Value: 2,
				Key:   KONSENTRASI,
			},
			LOGIKA_PRAKTIS: {
				Value: 3,
				Key:   LOGIKA_PRAKTIS,
			},
			FLEKSIBILITAS_BERPIKIR: {
				Value: 4,
				Key:   FLEKSIBILITAS_BERPIKIR,
			},
			IMAJINASI_KREATIF: {
				Value: 2,
				Key:   IMAJINASI_KREATIF,
			},
			ANTISIPASI: {
				Value: 4,
				Key:   ANTISIPASI,
			},
			POTENSI_KECERDASAN: {
				Value: 4,
				Key:   POTENSI_KECERDASAN,
			},
		},
	}
	candidateTable[SIKAP_KERJA][USMAN] = gap.CadidateInfo{
		Name: USMAN,
		Value: map[gap.SubAspec]gap.CadidateVal{
			ENERGI_PSIKIS: {
				Value: 4,
				Key:   ENERGI_PSIKIS,
			},
			KETELITIAN: {
				Value: 5,
				Key:   KETELITIAN,
			},
			KEHATI_HATIAN: {
				Value: 5,
				Key:   KEHATI_HATIAN,
			},
			PENGENDALIAN_PERASAAN: {
				Value: 1,
				Key:   PENGENDALIAN_PERASAAN,
			},
			DORONGAN_PRESTASI: {
				Value: 4,
				Key:   DORONGAN_PRESTASI,
			},
			VITALITAS_PERENCANAAN: {
				Value: 1,
				Key:   VITALITAS_PERENCANAAN,
			},
		},
	}
	candidateTable[PERILAKU][USMAN] = gap.CadidateInfo{
		Name: USMAN,
		Value: map[gap.SubAspec]gap.CadidateVal{
			DOMINACE: {
				Value: 4,
				Key:   DOMINACE,
			},
			INFLUENCES: {
				Value: 3,
				Key:   INFLUENCES,
			},
			STEADINESS: {
				Value: 4,
				Key:   STEADINESS,
			},
			COMPLIANCE: {
				Value: 4,
				Key:   COMPLIANCE,
			},
		},
	}
	//Kevin
	const TANTRI = "Tantri"
	candidateTable[KECERDASAN][TANTRI] = gap.CadidateInfo{
		Name: TANTRI,
		Value: map[gap.SubAspec]gap.CadidateVal{
			COMMON_SENSE: {
				Value: 3,
				Key:   COMMON_SENSE,
			},
			VERBALISASI_IDE: {
				Value: 5,
				Key:   VERBALISASI_IDE,
			},
			SISTEMATIKA_BERPIKIR: {
				Value: 4,
				Key:   SISTEMATIKA_BERPIKIR,
			},
			PENALARAN_DAN_SOLUSI_REAL: {
				Value: 3,
				Key:   PENALARAN_DAN_SOLUSI_REAL,
			},
			KONSENTRASI: {
				Value: 4,
				Key:   KONSENTRASI,
			},
			LOGIKA_PRAKTIS: {
				Value: 4,
				Key:   LOGIKA_PRAKTIS,
			},
			FLEKSIBILITAS_BERPIKIR: {
				Value: 3,
				Key:   FLEKSIBILITAS_BERPIKIR,
			},
			IMAJINASI_KREATIF: {
				Value: 5,
				Key:   IMAJINASI_KREATIF,
			},
			ANTISIPASI: {
				Value: 4,
				Key:   ANTISIPASI,
			},
			POTENSI_KECERDASAN: {
				Value: 3,
				Key:   POTENSI_KECERDASAN,
			},
		},
	}
	candidateTable[SIKAP_KERJA][TANTRI] = gap.CadidateInfo{
		Name: TANTRI,
		Value: map[gap.SubAspec]gap.CadidateVal{
			ENERGI_PSIKIS: {
				Value: 1,
				Key:   ENERGI_PSIKIS,
			},
			KETELITIAN: {
				Value: 5,
				Key:   KETELITIAN,
			},
			KEHATI_HATIAN: {
				Value: 5,
				Key:   KEHATI_HATIAN,
			},
			PENGENDALIAN_PERASAAN: {
				Value: 5,
				Key:   PENGENDALIAN_PERASAAN,
			},
			DORONGAN_PRESTASI: {
				Value: 5,
				Key:   DORONGAN_PRESTASI,
			},
			VITALITAS_PERENCANAAN: {
				Value: 2,
				Key:   VITALITAS_PERENCANAAN,
			},
		},
	}
	candidateTable[PERILAKU][TANTRI] = gap.CadidateInfo{
		Name: TANTRI,
		Value: map[gap.SubAspec]gap.CadidateVal{
			DOMINACE: {
				Value: 3,
				Key:   DOMINACE,
			},
			INFLUENCES: {
				Value: 3,
				Key:   INFLUENCES,
			},
			STEADINESS: {
				Value: 4,
				Key:   STEADINESS,
			},
			COMPLIANCE: {
				Value: 5,
				Key:   COMPLIANCE,
			},
		},
	}
	//Enrico
	const JAMES = "James"
	candidateTable[KECERDASAN][JAMES] = gap.CadidateInfo{
		Name: JAMES,
		Value: map[gap.SubAspec]gap.CadidateVal{
			COMMON_SENSE: {
				Value: 3,
				Key:   COMMON_SENSE,
			},
			VERBALISASI_IDE: {
				Value: 3,
				Key:   VERBALISASI_IDE,
			},
			SISTEMATIKA_BERPIKIR: {
				Value: 3,
				Key:   SISTEMATIKA_BERPIKIR,
			},
			PENALARAN_DAN_SOLUSI_REAL: {
				Value: 1,
				Key:   PENALARAN_DAN_SOLUSI_REAL,
			},
			KONSENTRASI: {
				Value: 2,
				Key:   KONSENTRASI,
			},
			LOGIKA_PRAKTIS: {
				Value: 5,
				Key:   LOGIKA_PRAKTIS,
			},
			FLEKSIBILITAS_BERPIKIR: {
				Value: 3,
				Key:   FLEKSIBILITAS_BERPIKIR,
			},
			IMAJINASI_KREATIF: {
				Value: 2,
				Key:   IMAJINASI_KREATIF,
			},
			ANTISIPASI: {
				Value: 5,
				Key:   ANTISIPASI,
			},
			POTENSI_KECERDASAN: {
				Value: 4,
				Key:   POTENSI_KECERDASAN,
			},
		},
	}
	candidateTable[SIKAP_KERJA][JAMES] = gap.CadidateInfo{
		Name: JAMES,
		Value: map[gap.SubAspec]gap.CadidateVal{
			ENERGI_PSIKIS: {
				Value: 4,
				Key:   ENERGI_PSIKIS,
			},
			KETELITIAN: {
				Value: 5,
				Key:   KETELITIAN,
			},
			KEHATI_HATIAN: {
				Value: 4,
				Key:   KEHATI_HATIAN,
			},
			PENGENDALIAN_PERASAAN: {
				Value: 3,
				Key:   PENGENDALIAN_PERASAAN,
			},
			DORONGAN_PRESTASI: {
				Value: 5,
				Key:   DORONGAN_PRESTASI,
			},
			VITALITAS_PERENCANAAN: {
				Value: 3,
				Key:   VITALITAS_PERENCANAAN,
			},
		},
	}
	candidateTable[PERILAKU][JAMES] = gap.CadidateInfo{
		Name: JAMES,
		Value: map[gap.SubAspec]gap.CadidateVal{
			DOMINACE: {
				Value: 4,
				Key:   DOMINACE,
			},
			INFLUENCES: {
				Value: 3,
				Key:   INFLUENCES,
			},
			STEADINESS: {
				Value: 3,
				Key:   STEADINESS,
			},
			COMPLIANCE: {
				Value: 5,
				Key:   COMPLIANCE,
			},
		},
	}

	//table Gap Weight Value

	wTable := gap.WeigTable{
		"0": gap.InfoWeig{
			Value:       5,
			Description: "Tidak ada selisih (kompetensi sesuai dgn yg dibutuhkan)",
		},
		"1": gap.InfoWeig{
			Value:       4.5,
			Description: "Kompetensi individu kelebihan 1 tingkat",
		},
		"-1": gap.InfoWeig{
			Value:       4,
			Description: "Kompetensi individu kekurangan 1 tingkat",
		},
		"2": gap.InfoWeig{
			Value:       3.5,
			Description: "Kompetensi individu kelebihan 2 tingkat",
		},
		"-2": gap.InfoWeig{
			Value:       3,
			Description: "Kompetensi individu kekurangan 2 tingkat",
		},
		"3": gap.InfoWeig{
			Value:       2.5,
			Description: "Kompetensi individu kelebihan 3 tingkat",
		},
		"-3": gap.InfoWeig{
			Value:       2,
			Description: "Kompetensi individu kekurangan 3 tingkat",
		},
		"4": gap.InfoWeig{
			Value:       1.5,
			Description: "Kompetensi individu kelebihan 4 tingkat",
		},
		"-4": gap.InfoWeig{
			Value:       1,
			Description: "Kompetensi individu kekurangan 4 tingkat",
		},
	}

	specT := gap.SpecTable{}
	specT.AspectTable = aspectTable
	specT.CadidateTable = candidateTable
	specT.Wtable = wTable

	result := CreacteGap(specT)

	return result
}
