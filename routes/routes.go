package routes

import (
	"UNISA_Server/controllers"
	"UNISA_Server/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin/register", controllers.Register)
	router.HandleFunc("/admin/login", controllers.Login)
	// ListTautan routes
	router.HandleFunc("/admin/list_tautan/create", middleware.AuthMiddleware(controllers.CreateListTautan)).Methods("POST")
	router.HandleFunc("/admin/list_tautan", middleware.AuthMiddleware(controllers.GetAllListTautan)).Methods("GET")
	router.HandleFunc("/admin/list_tautan/get", middleware.AuthMiddleware(controllers.GetListTautan)).Methods("GET")
	router.HandleFunc("/admin/list_tautan/update", middleware.AuthMiddleware(middleware.AdminMiddleware(controllers.UpdateListTautan))).Methods("PUT")
	router.HandleFunc("/admin/list_tautan/delete", middleware.AuthMiddleware(middleware.AdminMiddleware(controllers.DeleteListTautan))).Methods("DELETE")

	// DataLeads routes
	router.HandleFunc("/admin/data_leads/create", middleware.AuthMiddleware(controllers.CreateDataLeads)).Methods("POST")
	router.HandleFunc("/admin/data_leads", middleware.AuthMiddleware(controllers.GetAllDataLeads)).Methods("GET")
	router.HandleFunc("/admin/data_leads/get", middleware.AuthMiddleware(controllers.GetDataLeads)).Methods("GET")
	router.HandleFunc("/admin/data_leads/update", middleware.AuthMiddleware(middleware.AdminMiddleware(controllers.UpdateDataLeads))).Methods("PUT")
	router.HandleFunc("/admin/data_leads/delete", middleware.AuthMiddleware(middleware.AdminMiddleware(controllers.DeleteDataLeads))).Methods("DELETE")

	router.HandleFunc("/mahasiswa/register-mahasiswa", controllers.RegisterMahasiswa).Methods("POST")
	router.HandleFunc("/mahasiswa/presensi-mahasiswa", controllers.CreatePresensiMahasiswa).Methods("POST")
	// Testing route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is working!"))
	}).Methods("GET")
}
