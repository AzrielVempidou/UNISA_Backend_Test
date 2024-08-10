// controllers/controller.go
package controllers

import (
	"UNISA_Server/config"
	"UNISA_Server/models"
	"UNISA_Server/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterMahasiswa handles the registration of a new student
func RegisterMahasiswa(w http.ResponseWriter, r *http.Request) {
    var mahasiswa models.Mahasiswa
    err := json.NewDecoder(r.Body).Decode(&mahasiswa)
    if err != nil {
        utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
        return
    }

    result, err := config.MongoClient.Database("mydatabase").Collection("mahasiswa").InsertOne(context.TODO(), mahasiswa)
    if err != nil {
        utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating Mahasiswa")
        return
    }

    // Retrieve the newly created document to include all fields in the response
    var createdMahasiswa models.Mahasiswa
    err = config.MongoClient.Database("mydatabase").Collection("mahasiswa").FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&createdMahasiswa)
    if err != nil {
        utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching created Mahasiswa")
        return
    }

    response := map[string]interface{}{
        "status":      true,
        "Status_Code": http.StatusCreated,
        "message":     "Create Success",
        "data": map[string]interface{}{
            "ID":               createdMahasiswa.ID.Hex(),
            "NamaLengkap":      createdMahasiswa.NamaLengkap,
            "Alamat":           createdMahasiswa.Alamat,
            "TempatLahir":      createdMahasiswa.TempatLahir,
            "TanggalLahir":     createdMahasiswa.TanggalLahir,
            "Email":            createdMahasiswa.Email,
            "NoHp":             createdMahasiswa.NoHp,
            "NamaPTInstansi":   createdMahasiswa.NamaPTInstansi,
            "Jabatan":          createdMahasiswa.Jabatan,
						"NamaInstagram": 		createdMahasiswa.NamaInstagram,
						"NamaFacebook": 		createdMahasiswa.NamaFacebook,
						"SumberInformasi": 	createdMahasiswa.SumberInformasi,
        },
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// CreatePresensiMahasiswa handles the creation of a student attendance record
func CreatePresensiMahasiswa(w http.ResponseWriter, r *http.Request) {
	var mahasiswa models.Mahasiswa
	var presensi models.Presensi // Presensi model should be defined with `TanggalPresensi` and `Status` fields.

	// Decode the request body into the Mahasiswa model
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	mahasiswaCollection := config.MongoClient.Database("mydatabase").Collection("mahasiswa")
	filter := bson.M{
		"nama_lengkap": mahasiswa.NamaLengkap,
		"alamat":       mahasiswa.Alamat,
		"tanggal_lahir": mahasiswa.TanggalLahir,
		"email":        mahasiswa.Email,
		"no_hp":        mahasiswa.NoHp,
		"nama_pt_instansi": mahasiswa.NamaPTInstansi,
		"jabatan":      mahasiswa.Jabatan,
		"nama_instagram": mahasiswa.NamaInstagram,
		"nama_facebook":  mahasiswa.NamaFacebook,
	}

	var existingMahasiswa models.Mahasiswa
	err = mahasiswaCollection.FindOne(context.TODO(), filter).Decode(&existingMahasiswa)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.ErrorResponse(w, http.StatusNotFound, "Mahasiswa not found")
		} else {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching Mahasiswa")
		}
		return
	}

	presensi.TanggalPresensi = time.Now().Format("2006-01-02") // Format date to YYYY-MM-DD
	presensi.Status = "Hadir" // Assuming default status is "Hadir"

	presensiCollection := config.MongoClient.Database("mydatabase").Collection("presensi")
	_, err = presensiCollection.InsertOne(context.TODO(), presensi)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating presensi record")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusCreated,
		"message":     "Presensi Success",
		"data": map[string]interface{}{
			"nama_lengkap":     existingMahasiswa.NamaLengkap,
			"alamat":          existingMahasiswa.Alamat,
			"tanggal_lahir":   existingMahasiswa.TanggalLahir,
			"email":           existingMahasiswa.Email,
			"no_hp":           existingMahasiswa.NoHp,
			"nama_pt_instansi": existingMahasiswa.NamaPTInstansi,
			"jabatan":         existingMahasiswa.Jabatan,
			"nama_instagram":  existingMahasiswa.NamaInstagram,
			"nama_facebook":   existingMahasiswa.NamaFacebook,
			"tanggal_presensi": presensi.TanggalPresensi,
			"status":          presensi.Status,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
