package controllers

import (
	"fmt"
	"log"
	"net/http"
	"teste-egsys/server/models"
	"teste-egsys/server/responses"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Fleet{}, &models.Vehicle{}, &models.Fleet_Alert{}, &models.Vehicle_Position{}) //database migration
	server.DB.Debug().Model(&models.Fleet_Alert{}).AddForeignKey("fleet_id", "fleets(fleet_id)", "cascade", "cascade")
	server.DB.Debug().Model(&models.Vehicle{}).AddForeignKey("fleet_id", "fleets(fleet_id)", "cascade", "cascade")
	server.DB.Debug().Model(&models.Vehicle_Position{}).AddForeignKey("vehicle_id", "vehicles(vehicle_id)", "cascade", "cascade")
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Reset(w http.ResponseWriter, r *http.Request) {

	Fleet_Alert := models.Fleet_Alert{}

	_, err := Fleet_Alert.DeleteAllFleetAlert(server.DB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Vehicle_Position := models.Vehicle_Position{}

	_, err = Vehicle_Position.DeleteAllVehiclePosition(server.DB)
	if err != nil {
		return
	}

	vehicle := models.Vehicle{}

	_, err = vehicle.DeleteAllVehicle(server.DB)
	if err != nil {
		return
	}

	fleet := models.Fleet{}

	_, err = fleet.DeleteAllFleet(server.DB)
	if err != nil {
		return
	}

	responses.JSON(w, http.StatusOK, "Base resetado com sucesso")
}
