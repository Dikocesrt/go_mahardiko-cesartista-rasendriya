package model

type Package struct {
	Id             int     `json:"name"`
	Nama           string  `json:"sender"`
	NamaPengirim   string  `json:"receiver"`
	NamaPenerima   string  `json:"sender_location"`
	LokasiPengirim string  `json:"sender_receiver"`
	Biaya          int     `json:"fee"`
	BeratPaket     float64 `json:"weight"`
}