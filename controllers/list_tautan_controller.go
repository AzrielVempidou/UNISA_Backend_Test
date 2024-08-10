package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"UNISA_Server/config"
	"UNISA_Server/models"
	"UNISA_Server/utils"

	"github.com/boombuler/barcode/qr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateListTautan menangani pembuatan ListTautan dengan pembuatan kode QR
func CreateListTautan(w http.ResponseWriter, r *http.Request) {
	var listTautan models.ListTautan
	err := json.NewDecoder(r.Body).Decode(&listTautan)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	qrContent := fmt.Sprintf("Nama Program: %s, Nama Instansi: %s, Nama Kegiatan: %s, Alamat: %s, Nama PIC: %s, Nama PIC PT/Instansi: %s, Tanggal Mulai: %s, Tanggal Akhir: %s",
		listTautan.NamaProgram, listTautan.NamaInstansi, listTautan.NamaKegiatan, listTautan.Alamat, listTautan.NamaPIC, listTautan.NamaPICPTInstansi, listTautan.TanggalMulai, listTautan.TanggalAkhir)

	barcode, err := qr.Encode(qrContent, qr.L, qr.Auto)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating QR code")
		return
	}

	objectID := primitive.NewObjectID()
	listTautan.ID = objectID
	fileName := fmt.Sprintf("%s.png", objectID.Hex())
	filePath := filepath.Join("qrcodes", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error saving QR code")
		return
	}
	defer file.Close()

	png.Encode(file, barcode)

	listTautan.QRCodePath = filePath

	_, err = config.MongoClient.Database("mydatabase").Collection("list_tautan").InsertOne(context.TODO(), listTautan)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating ListTautan")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusCreated,
		"message":     "Create Success",
		"data": map[string]interface{}{
			"ID":                listTautan.ID.Hex(),
			"QRCodePath":        listTautan.QRCodePath,
			"NamaProgram":       listTautan.NamaProgram,
			"NamaInstansi":      listTautan.NamaInstansi,
			"NamaKegiatan":      listTautan.NamaKegiatan,
			"Alamat":            listTautan.Alamat,
			"NamaPIC":           listTautan.NamaPIC,
			"NamaPICPTInstansi": listTautan.NamaPICPTInstansi,
			"TanggalMulai":      listTautan.TanggalMulai,
			"TanggalAkhir":      listTautan.TanggalAkhir,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func GetAllListTautan(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")

	page := 1
	limit := 10
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid page number")
			return
		}
	}
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid limit number")
			return
		}
	}
	if limit > 100 {
		limit = 100 // Limit maximum number of items per page
	}

	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"NamaProgram": bson.M{"$regex": search, "$options": "i"}},
				{"NamaInstansi": bson.M{"$regex": search, "$options": "i"}},
				{"NamaKegiatan": bson.M{"$regex": search, "$options": "i"}},
				{"Alamat": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	sort := bson.M{}
	if sortBy != "" {
		order := 1 
		if sortOrder == "desc" {
			order = -1
		}
		sort[sortBy] = order
	}

	collection := config.MongoClient.Database("mydatabase").Collection("list_tautan")

	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  int64PointerListTautan(int64((page - 1) * limit)),
		Limit: int64PointerListTautan(int64(limit)),
		Sort:  sort,
	})
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.TODO())

	var listTautan []models.ListTautan
	if err = cursor.All(context.TODO(), &listTautan); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var responseData []map[string]interface{}
	for _, item := range listTautan {
		dataItem := map[string]interface{}{
			"ID":                item.ID.Hex(),
			"QRCodePath":        item.QRCodePath,
			"NamaProgram":       item.NamaProgram,
			"NamaInstansi":      item.NamaInstansi,
			"NamaKegiatan":      item.NamaKegiatan,
			"Alamat":            item.Alamat,
			"NamaPIC":           item.NamaPIC,
			"NamaPICPTInstansi": item.NamaPICPTInstansi,
			"TanggalMulai":      item.TanggalMulai,
			"TanggalAkhir":      item.TanggalAkhir,
		}
		responseData = append(responseData, dataItem)
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": 200,
		"message":     "Data ditemukan",
		"data":        responseData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Helper function to convert int64 to *int64
func int64PointerListTautan(i int64) *int64 {
	return &i
}


// GetListTautan menangani pengambilan ListTautan berdasarkan ID
func GetListTautan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Query tidak boleh kosong")
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var listTautan models.ListTautan
	err = config.MongoClient.Database("mydatabase").Collection("list_tautan").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&listTautan)
	if err == mongo.ErrNoDocuments {
		utils.ErrorResponse(w, http.StatusNotFound, "Data tidak ditemukan")
		return
	} else if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching ListTautan")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Data ditemukan",
		"data": map[string]interface{}{
			"ID":                listTautan.ID.Hex(),
			"QRCodePath":        listTautan.QRCodePath,
			"NamaProgram":       listTautan.NamaProgram,
			"NamaInstansi":      listTautan.NamaInstansi,
			"NamaKegiatan":      listTautan.NamaKegiatan,
			"Alamat":            listTautan.Alamat,
			"NamaPIC":           listTautan.NamaPIC,
			"NamaPICPTInstansi": listTautan.NamaPICPTInstansi,
			"TanggalMulai":      listTautan.TanggalMulai,
			"TanggalAkhir":      listTautan.TanggalAkhir,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


// UpdateListTautan menangani pembaruan ListTautan berdasarkan ID
func UpdateListTautan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Query tidak boleh kosong")
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var listTautan models.ListTautan
	err = json.NewDecoder(r.Body).Decode(&listTautan)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	var existingListTautan models.ListTautan
	err = config.MongoClient.Database("mydatabase").Collection("list_tautan").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&existingListTautan)
	if err == mongo.ErrNoDocuments {
		utils.ErrorResponse(w, http.StatusNotFound, "Data tidak ditemukan")
		return
	} else if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching ListTautan")
		return
	}

	if existingListTautan.QRCodePath != "" {
		err = os.Remove(existingListTautan.QRCodePath)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error removing old QR code")
			return
		}
	}

	qrContent := fmt.Sprintf("Nama Program: %s, Nama Instansi: %s, Nama Kegiatan: %s, Alamat: %s, Nama PIC: %s, Nama PIC PT/Instansi: %s, Tanggal Mulai: %s, Tanggal Akhir: %s",
		listTautan.NamaProgram, listTautan.NamaInstansi, listTautan.NamaKegiatan, listTautan.Alamat, listTautan.NamaPIC, listTautan.NamaPICPTInstansi, listTautan.TanggalMulai, listTautan.TanggalAkhir)

	barcode, err := qr.Encode(qrContent, qr.L, qr.Auto)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating QR code")
		return
	}

	fileName := fmt.Sprintf("%s.png", objID.Hex())
	filePath := filepath.Join("qrcodes", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error saving QR code")
		return
	}
	defer file.Close()

	png.Encode(file, barcode)

	listTautan.QRCodePath = filePath
	update := bson.M{"$set": listTautan}
	_, err = config.MongoClient.Database("mydatabase").Collection("list_tautan").UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error updating ListTautan")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusCreated,
		"message":     "Update Success",
		"data": map[string]interface{}{
			"ID":                objID.Hex(),
			"QRCodePath":        listTautan.QRCodePath,
			"NamaProgram":       listTautan.NamaProgram,
			"NamaInstansi":      listTautan.NamaInstansi,
			"NamaKegiatan":      listTautan.NamaKegiatan,
			"Alamat":            listTautan.Alamat,
			"NamaPIC":           listTautan.NamaPIC,
			"NamaPICPTInstansi": listTautan.NamaPICPTInstansi,
			"TanggalMulai":      listTautan.TanggalMulai,
			"TanggalAkhir":      listTautan.TanggalAkhir,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func DeleteListTautan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Query tidak boleh kosong")
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var listTautan models.ListTautan
	err = config.MongoClient.Database("mydatabase").Collection("list_tautan").FindOneAndDelete(context.TODO(), bson.M{"_id": objID}).Decode(&listTautan)
	if err == mongo.ErrNoDocuments {
		utils.ErrorResponse(w, http.StatusNotFound, "Data tidak ditemukan")
		return
	} else if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error deleting ListTautan")
		return
	}

	if listTautan.QRCodePath != "" {
		err = os.Remove(listTautan.QRCodePath)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error removing QR code")
			return
		}
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Delete List Tautan Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


