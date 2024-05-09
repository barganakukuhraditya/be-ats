package be_ats_test

import (
	"fmt"
	"testing"

	"github.com/barganakukuhraditya/be_ats/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertJadwal(t *testing.T) {
	mataKuliah := map[string]interface{}{
		"_id":          primitive.NewObjectID(),
		"kode_matkul": "MK101",
		"nama_matkul": "Matematika Dasar",
		"sks":         3,
		"semester":    1,
	}

	dosen := map[string]interface{}{
		"nidn":       "D101",
		"nama_dosen": "Dosen 1",
	}

	insertedID := be_ats.InsertJadwal(
		"Senin",
		"08:00",
		"10:00",
		mataKuliah,
		101,
		dosen,
	)
	fmt.Println("Inserted Jadwal ID:", insertedID)
}

func TestGetAllJadwal(t *testing.T) {
	jadwals, err := be_ats.GetAllJadwal()
	if err != nil {
		t.Errorf("Error getting all jadwal: %v", err)
		return
	}

	fmt.Println("All Jadwal:")
	for _, jadwal := range jadwals {
		fmt.Println(jadwal)
	}
}
