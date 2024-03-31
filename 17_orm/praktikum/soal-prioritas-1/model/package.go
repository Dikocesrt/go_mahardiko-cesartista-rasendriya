package model

type Package struct {
	Id             int     `json:"id"`
	Nama           string  `json:"name"`
	NamaPengirim   string  `json:"sender"`
	NamaPenerima   string  `json:"receiver"`
	LokasiPengirim string  `json:"sender_location"`
	Biaya          int     `json:"fee"`
	BeratPaket     float64 `json:"weight"`
}