package be_ats

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect berfungsi untuk terhubung ke database MongoDB
func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// InsertOneDoc berfungsi untuk menyisipkan satu dokumen ke dalam koleksi MongoDB
func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// InsertJadwal berfungsi untuk menyisipkan jadwal ke dalam koleksi MongoDB
func InsertJadwal(hari string, waktuMulai string, waktuSelesai string, mataKuliah interface{}, ruangan int, dosen interface{}) interface{} {
	jadwal := map[string]interface{}{
		"_id":           primitive.NewObjectID(),
		"hari":          hari,
		"waktuMulai":    waktuMulai,
		"waktuSelesai":  waktuSelesai,
		"mataKuliah":    mataKuliah,
		"ruangan":       ruangan,
		"dosen":         dosen,
	}
	return InsertOneDoc("week4", "jadwal", jadwal)
}

// GetAllJadwal berfungsi untuk mengambil semua jadwal dari koleksi MongoDB
func GetAllJadwal() ([]interface{}, error) {
	collection := MongoConnect("week4").Collection("jadwal")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllJadwal:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var jadwal []interface{}
	if err := cursor.All(context.TODO(), &jadwal); err != nil {
		fmt.Println("GetAllJadwal:", err)
		return nil, err
	}
	return jadwal, nil
}
