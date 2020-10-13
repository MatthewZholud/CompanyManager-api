package main

import (
	"log"
	"net/http"

	"github.com/MatthewZholud/CompanyManager-api/internal/domain/routes"

	"github.com/gorilla/mux"

)

const (
	apiGatewayPort           = ":3001"
	employeeMicroServiceAddr = "localhost:3443"
	companyMicroServiceAddr  = "localhost:4443"
)

type someInterface interface {}

func main() {

	r := mux.NewRouter()

	routes.RegisterEmployeeRoutes(r)
	routes.RegisterCompanyRoutes(r)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))

}
