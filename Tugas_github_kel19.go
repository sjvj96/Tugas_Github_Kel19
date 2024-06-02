package main

import (
	"fmt"
	"strings"
)

type Kendaraan struct {
	Nama  string
	Plat  string
	Tipe  string
	Rute  []string
	Tilang int
}

func kenaRazia(tanggal int, data []Kendaraan) []Kendaraan {
	var pelanggaran []Kendaraan

	for _, kendaraan := range data {
		platAngka := kendaraan.Plat[len(kendaraan.Plat)-1]

		if tanggal >= 1 && tanggal <= 31 {
			if platAngka%2 == 0 {
				for _, lokasi := range kendaraan.Rute {
					if strings.Contains(lokasi, "Gajah Mada") ||
						strings.Contains(lokasi, "Hayam Wuruk") ||
						strings.Contains(lokasi, "Sisingamangaraja") ||
						strings.Contains(lokasi, "Panglima Polim") ||
						strings.Contains(lokasi, "Fatmawati") ||
						strings.Contains(lokasi, "Tomang Raya") {
						kendaraan.Tilang++
					}
				}
			}
		}

		pelanggaran = append(pelanggaran, kendaraan)
	}

	return pelanggaran
}

func main() {
	dataKendaraan := []Kendaraan{
		{
			Nama:  "Denver",
			Plat:  "B 2791 KDS",
			Tipe:  "Mobil",
			Rute:  []string{"TB Simatupang", "Panglima Polim", "Depok", "Senen Raya"},
		},
		{
			Nama:  "Toni",
			Plat:  "B 1212 JBB",
			Tipe:  "Mobil",
			Rute:  []string{"Pintu Besar Selatan", "Panglima Polim", "Depok", "Senen Raya", "Kemang"},
		},
		{
			Nama:  "Stark",
			Plat:  "B 444 XSX",
			Tipe:  "Motor",
			Rute:  []string{"Pondok Indah", "Depok", "Senen Raya", "Kemang"},
		},
		{
			Nama:  "Anna",
			Plat:  "B 678 DD",
			Tipe:  "Mobil",
			Rute:  []string{"Fatmawati", "Panglima Polim", "Depok", "Senen Raya", "Kemang", "Gajah Mada"},
		},
	}

	hasil := kenaRazia(27, dataKendaraan)
	for _, kendaraan := range hasil {
		fmt.Printf("%s %s %s %v\n", kendaraan.Nama, kendaraan.Plat, kendaraan.Tipe, kendaraan.Tilang)
	}
}
