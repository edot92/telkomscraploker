package db

import "log"

type DataFb struct {
	Waktu string `bson:"waktu"`
	Post  string `bson:"post"`
}

// StructDataFb ...
type StructDataFb struct {
	URL    string `bson:"url"`
	DataFb []struct {
		Waktu string `bson:"waktu"`
		Post  string `bson:"post"`
	} `bson:"dataFb"`
}

// InsertDatafb ...
func InsertDatafb(data StructDataFb) {
	// SessionMgo.
	sesMgo := DBCopyMGO()
	defer sesMgo.Close()
	// Collection People
	col := sesMgo.DB(DBNAME).C(ColFb)
	err := col.Insert(data)
	if err != nil {
		panic(err)
	}
	log.Println("succes insert")
}
