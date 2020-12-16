package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xans-me/protobf-grpc/common/config"
	"github.com/xans-me/protobf-grpc/common/model"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "u0001",
		Name:     "Mulia Ichsan",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	// garage1 := model.Garage{
	// 	Id:   "q001",
	// 	Name: "Quel'thalas",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  45.123123123,
	// 		Longitude: 54.1231313123,
	// 	},
	// }

	userSvc := serviceUser()

	fmt.Println("\n", "===========> user test")

	userSvc.Register(context.Background(), &user1)
}
