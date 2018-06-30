package crawl

import (
	"encoding/csv"
	"log"
	"os"
	"sample/db"
	"strings"
)

type StructCsv struct { // Our example struct, you can use "-" to ignore a field
	DataLowongan string `csv:"dataLowongan"`
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Parse1 ...
func Parse1(txt string) {
	// log.Fatal(txt)
	txt = strings.Replace(txt, "Bahasa Indonesia · English (UK) · Basa Jawa · Español · Português (Brasil)Privasi · Ketentuan · Iklan · Pilihan Iklan · Cookie · Lainnya Facebook © 2018   Kabar BeritaAKTIVITAS TERBARU", "", -1)
	txt = strings.Replace(txt, "Lihat Selengkapnya", "", -1)
	dataSPlit1 := strings.Split(txt, "SukaKomentariBagikan")
	if len(dataSPlit1) == 0 {
		log.Fatal("tidak ada data untuk split")
	}
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	tanggalKemarin := ""
	var listDataKemarenOrJamSaja []string
	var dataInsert db.StructDataFb
	// var tempListDataFb []db.DataFb
	for _, value := range dataSPlit1 {
		if strings.Contains(value, " menyukai ini.") {
			value2 := strings.Split(value, " menyukai ini.")
			if len(value2) > 1 {
				value = value2[1]
			}
		}
		value2 := strings.Split(value, " ·")
		if len(value2) > 1 {
			value = ""
			for index := 0; index < len(value2); index++ {
				if index < len(value2)-1 {
					value = value + value2[index+1]
				}
			}
		}
		// remoe string yang tifak perlu
		value = strings.Replace(value, "menonaktifkan komentar untuk kiriman ini.", "", -1)
		value = strings.Replace(value, "Kiriman Terdahulu", "", -1)
		// parsing pukul
		// 	 24 Juni pukul 21.00
		if strings.Contains(value, " pukul ") && strings.Contains(value, "Kemarin pukul") == false {
			value2 := strings.Split(value, "pukul")
			if len(value2) > 1 {
				value = ""
				for index := 0; index < len(value2); index++ {
					if index < len(value2)-1 {
						value = value + value2[index+1]
					}
				}
				// remove jam
				valueTemp := value
				value = ""
				for index := 0; index < len(valueTemp); index++ {
					if index > 5 {
						value = value + string(valueTemp[index])
					}
				}
				dataInsert.DataFb = append(dataInsert.DataFb, db.DataFb{
					Waktu: value2[0],
					Post:  value,
				})
			}
			if strings.Contains(value, "Kemarin") {
				log.Println("tanggal 1 :" + tanggalKemarin)
				listDataKemarenOrJamSaja = append(listDataKemarenOrJamSaja, value)
			} else {
				if tanggalKemarin == "" {
					tanggalKemarin = value2[0]
					listDataKemarenOrJamSaja = append(listDataKemarenOrJamSaja, value)
				}
				log.Println("tanggal 2 :" + value2[0])
			}
			// tempListDataFb = append(tempListDataFb, db.DataFb{
			// 	Waktu: value2[0],
			// 	Post:  value,
			// })

		} else if strings.Contains(value, "Kemarin pukul") {
			log.Println("tanggal 3 :" + tanggalKemarin)
			listDataKemarenOrJamSaja = append(listDataKemarenOrJamSaja, value)
		} else {
			if strings.Contains(value, "jam") {
				log.Println("tanggal 4 :" + tanggalKemarin)
			}
		}
		if tanggalKemarin != "" {
			log.Println("DATA POST:" + value)
			log.Println("-------*****-------")
		}

		// 	newVal := []string{
		// 		value,
		// 	}
		// 	err := writer.Write(newVal)
		// 	checkError("Cannot write to file", err)
	}
	// insert array to database

	dataInsert.URL = "tse"
	// TODO proses tanggal kemaren dan jam
	// if len(listDataKemarenOrJamSaja) > 0 {
	// 	for _, value := range listDataKemarenOrJamSaja {
	// 		log.Println(value)
	// 		// parsing pukul
	// 		// 	 24 Juni pukul 21.00
	// 		if strings.Contains(value, "Kemarin pukul") || strings.Contains(value, "Jam") {
	// 			value2 := strings.Split(value, "Kemarin pukul")
	// 			if len(value2) > 1 {
	// 				value = ""
	// 				for index := 0; index < len(value2); index++ {
	// 					if index < len(value2)-1 {
	// 						value = value + value2[index+1]
	// 					}
	// 				}
	// 				// remove jam
	// 				valueTemp := value
	// 				value = ""
	// 				for index := 0; index < len(valueTemp); index++ {
	// 					if index > 5 {
	// 						value = value + string(valueTemp[index])
	// 					}
	// 				}
	// 			}
	// 			log.Println("tanggal kemaren :" + tanggalKemarin)
	// 			log.Println("DATA POST:" + value)
	// 			log.Println("-------*****-------")
	// 		}

	// 	}
	// }
	db.InsertDatafb(dataInsert)
	_ = listDataKemarenOrJamSaja
}
