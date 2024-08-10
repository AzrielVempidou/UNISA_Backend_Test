package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"UNISA_Server/config"
	"UNISA_Server/models"
	"UNISA_Server/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateDataLeads handles the creation of DataLeads
func CreateDataLeads(w http.ResponseWriter, r *http.Request) {
	var dataLeads models.DataLeads
	err := json.NewDecoder(r.Body).Decode(&dataLeads)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	result, err := config.MongoClient.Database("mydatabase").Collection("data_leads").InsertOne(context.TODO(), dataLeads)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error creating DataLeads")
		return
	}

	var createdDataLeads models.DataLeads
	err = config.MongoClient.Database("mydatabase").Collection("data_leads").FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&createdDataLeads)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching created DataLeads")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusCreated,
		"message":     "Create Success",
		"data": map[string]interface{}{
			"ID":               createdDataLeads.ID.Hex(),
			"NamaLengkap":      createdDataLeads.NamaLengkap,
			"Alamat":           createdDataLeads.Alamat,
			"TempatLahir":      createdDataLeads.TempatLahir,
			"TanggalLahir":     createdDataLeads.TanggalLahir, // Directly use the string value
			"Email":            createdDataLeads.Email,
			"NoHp":             createdDataLeads.NoHp,
			"NamaPTInstansi":   createdDataLeads.NamaPTInstansi,
			"Jabatan":          createdDataLeads.Jabatan,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



// GetAllDataLeads handles fetching all DataLeads with pagination, search, and sorting
func GetAllDataLeads(w http.ResponseWriter, r *http.Request) {
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
				{"nama_lengkap": bson.M{"$regex": search, "$options": "i"}},
				{"alamat": bson.M{"$regex": search, "$options": "i"}},
				{"tempat_lahir": bson.M{"$regex": search, "$options": "i"}},
				{"email": bson.M{"$regex": search, "$options": "i"}},
				{"no_hp": bson.M{"$regex": search, "$options": "i"}},
				{"nama_pt_instansi": bson.M{"$regex": search, "$options": "i"}},
				{"jabatan": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	sort := bson.M{}
	if sortBy != "" {
		order := 1 // Default ascending
		if sortOrder == "desc" {
			order = -1
		}
		sort[sortBy] = order
	}

	collection := config.MongoClient.Database("mydatabase").Collection("data_leads")

	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  int64PointerDataLeads(int64((page - 1) * limit)),
		Limit: int64PointerDataLeads(int64(limit)),
		Sort:  sort,
	})
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching DataLeads")
		return
	}
	defer cursor.Close(context.TODO())

	var dataLeads []models.DataLeads
	if err = cursor.All(context.TODO(), &dataLeads); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching DataLeads")
		return
	}

	var responseData []map[string]interface{}
	for _, item := range dataLeads {
		dataItem := map[string]interface{}{
			"ID":               item.ID.Hex(),
			"NamaLengkap":      item.NamaLengkap,
			"Alamat":           item.Alamat,
			"TempatLahir":      item.TempatLahir,
			"TanggalLahir":     item.TanggalLahir, // Directly use the string value
			"Email":            item.Email,
			"NoHp":             item.NoHp,
			"NamaPTInstansi":   item.NamaPTInstansi,
			"Jabatan":          item.Jabatan,
		}
		responseData = append(responseData, dataItem)
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Data ditemukan",
		"data":        responseData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetDataLeads handles fetching a DataLeads by ID
func GetDataLeads(w http.ResponseWriter, r *http.Request) {
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

	var dataLeads models.DataLeads
	err = config.MongoClient.Database("mydatabase").Collection("data_leads").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&dataLeads)
	if err == mongo.ErrNoDocuments {
		utils.ErrorResponse(w, http.StatusNotFound, "Data tidak ditemukan")
		return
	} else if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching DataLeads")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Data ditemukan",
		"data": map[string]interface{}{
			"ID":               dataLeads.ID.Hex(),
			"NamaLengkap":      dataLeads.NamaLengkap,
			"Alamat":           dataLeads.Alamat,
			"TempatLahir":      dataLeads.TempatLahir,
			"TanggalLahir":     dataLeads.TanggalLahir, // Directly use the string value
			"Email":            dataLeads.Email,
			"NoHp":             dataLeads.NoHp,
			"NamaPTInstansi":   dataLeads.NamaPTInstansi,
			"Jabatan":          dataLeads.Jabatan,
			// Add other fields of DataLeads as necessary
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// UpdateDataLeads handles updating a DataLeads by ID
func UpdateDataLeads(w http.ResponseWriter, r *http.Request) {
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

	var dataLeads models.DataLeads
	err = json.NewDecoder(r.Body).Decode(&dataLeads)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	update := bson.M{"$set": dataLeads}
	_, err = config.MongoClient.Database("mydatabase").Collection("data_leads").UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error updating DataLeads")
		return
	}

	// Retrieve the updated document to include all fields in the response
	var updatedDataLeads models.DataLeads
	err = config.MongoClient.Database("mydatabase").Collection("data_leads").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&updatedDataLeads)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error fetching updated DataLeads")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Update Success",
		"data": map[string]interface{}{
			"ID":               updatedDataLeads.ID.Hex(),
			"NamaLengkap":      updatedDataLeads.NamaLengkap,
			"Alamat":           updatedDataLeads.Alamat,
			"TempatLahir":      updatedDataLeads.TempatLahir,
			"TanggalLahir":     updatedDataLeads.TanggalLahir, // Directly use the string value
			"Email":            updatedDataLeads.Email,
			"NoHp":             updatedDataLeads.NoHp,
			"NamaPTInstansi":   updatedDataLeads.NamaPTInstansi,
			"Jabatan":          updatedDataLeads.Jabatan,
			// Add other fields of DataLeads as necessary
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



// DeleteDataLeads handles deleting a DataLeads by ID
func DeleteDataLeads(w http.ResponseWriter, r *http.Request) {
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

	result, err := config.MongoClient.Database("mydatabase").Collection("data_leads").DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error deleting DataLeads")
		return
	}

	if result.DeletedCount == 0 {
		utils.ErrorResponse(w, http.StatusNotFound, "Data tidak ditemukan")
		return
	}

	response := map[string]interface{}{
		"status":      true,
		"Status_Code": http.StatusOK,
		"message":     "Delete Data Leads Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Helper function to convert int64 to *int64
func int64PointerDataLeads(i int64) *int64 {
	return &i
}