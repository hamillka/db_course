package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hamillka/ppo/backend/internal/handlers/middlewares"
	"go.uber.org/zap"
)

//nolint:funlen // it's ok
func Router(
	as AppointmentService,
	bs BranchService,
	ds DoctorService,
	os OfficeService,
	ps PatientService,
	us UserService,
	ts TimetableService,
	mhs MedicalHistoryService,
	logger *zap.SugaredLogger,
) *mux.Router {
	router := mux.NewRouter()
	secure := router.PathPrefix("/auth").Subrouter()
	fun := router.PathPrefix("/api/v1").Subrouter()

	fun.Use(middlewares.AuthMiddleware)

	ah := NewAppointmentHandler(as, logger)
	bh := NewBranchHandler(bs, logger)
	dh := NewDoctorHandler(ds, logger)
	oh := NewOfficeHandler(os, logger)
	ph := NewPatientHandler(ps, logger)
	uh := NewUserHandler(us, logger)
	th := NewTimetableHandler(ts, logger)
	mhh := NewMedicalHistoryHandler(mhs, logger)

	secure.HandleFunc("/login", uh.Login).Methods("POST")
	secure.HandleFunc("/register", uh.Register).Methods("POST")

	fun.HandleFunc("/appointment", ah.GetAppointments).Methods("GET")
	fun.HandleFunc("/appointment/{id}", ah.GetAppointmentByID).Methods("GET")
	fun.HandleFunc("/appointment/{id}", ah.CancelAppointment).Methods("DELETE")
	fun.HandleFunc("/appointment", ah.CreateAppointment).Methods("POST")
	fun.HandleFunc("/appointment/{id}", ah.EditAppointment).Methods("PATCH")

	fun.HandleFunc("/branch/{id}", bh.EditBranch).Methods("PATCH")
	fun.HandleFunc("/branch", bh.AddBranch).Methods("POST")
	fun.HandleFunc("/branch", bh.GetAllBranches).Methods("GET")
	fun.HandleFunc("/branch/{id}", bh.GetBranchByID).Methods("GET")

	fun.HandleFunc("/doctor/{id}", dh.EditDoctor).Methods("PATCH")
	fun.HandleFunc("/doctor", dh.AddDoctor).Methods("POST")
	fun.HandleFunc("/doctor", dh.GetAllDoctors).Methods("GET")
	fun.HandleFunc("/doctor/{id}", dh.GetDoctorByID).Methods("GET")

	fun.HandleFunc("/office/{id}", oh.EditOffice).Methods("PATCH")
	fun.HandleFunc("/office", oh.AddOffice).Methods("POST")
	fun.HandleFunc("/office", oh.GetAllOffices).Methods("GET")
	fun.HandleFunc("/office/{id}", oh.GetOfficeByID).Methods("GET")

	fun.HandleFunc("/patient/{id}", ph.EditPatient).Methods("PATCH")
	fun.HandleFunc("/patient", ph.AddPatient).Methods("POST")
	fun.HandleFunc("/patient", ph.GetAllPatients).Methods("GET")
	fun.HandleFunc("/patient/{id}", ph.GetPatientByID).Methods("GET")

	fun.HandleFunc("/loc_by_doc/{id}", th.GetLocationsByDoctor).Methods("GET")
	fun.HandleFunc("/doc_by_loc/{id}", th.GetDoctorsByLocation).Methods("GET")
	fun.HandleFunc("/workdays/{id}", th.GetTimetableByDoctor).Methods("GET")

	fun.HandleFunc("/medical_history/{id}", mhh.GetHistoryByPatient).Methods("GET")
	fun.HandleFunc("/medical_history", mhh.CreateMedicalHistory).Methods("POST")
	fun.HandleFunc("/medical_history/{id}", mhh.UpdateMedicalHistory).Methods("PATCH")

	return router
}
