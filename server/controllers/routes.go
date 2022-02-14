package controllers

import "teste-egsys/server/middlewares"

func (s *Server) initializeRoutes() {

	// Fleet route
	s.Router.HandleFunc("/fleet", middlewares.SetMiddleware(s.GetAllFleets)).Methods("GET")
	s.Router.HandleFunc("/fleet", middlewares.SetMiddleware(s.CreateFleet)).Methods("POST")

	// Vehicle route
	s.Router.HandleFunc("/vehicle", middlewares.SetMiddleware(s.GetAllVehicles)).Methods("GET")
	s.Router.HandleFunc("/vehicle", middlewares.SetMiddleware(s.CreateVehicle)).Methods("POST")

	// Fleet_Alert route
	s.Router.HandleFunc("/fleet/{id}/alert", middlewares.SetMiddleware(s.GetAllFleetAlerts)).Methods("GET")
	s.Router.HandleFunc("/fleet/{id}/alert", middlewares.SetMiddleware(s.CreateFleetAlert)).Methods("POST")

	// VehiclePosition route
	s.Router.HandleFunc("/vehicle/{id}/positon", middlewares.SetMiddleware(s.GetAllVehiclePositions)).Methods("GET")
	s.Router.HandleFunc("/vehicle/{id}/positon", middlewares.SetMiddleware(s.CreateVehiclePosition)).Methods("POST")

	// Reset
	s.Router.HandleFunc("/reset", middlewares.SetMiddleware(s.Reset)).Methods("DELETE")
}
