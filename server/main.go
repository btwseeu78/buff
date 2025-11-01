package main

import (
	"context"
	"fmt"
	"log"

	"net/http"

	petv1 "github.com/btwseeu78/buff/gen/pet/v1"
	"github.com/btwseeu78/buff/gen/pet/v1/petv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

type PetStoreServiceServer struct {
	petv1connect.UnimplementedPetStoreServiceHandler
}

func main() {
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(&PetStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("..listening on", address)
	http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *PetStoreServiceServer) PutPet(
	_ context.Context,
	req *petv1.PutPetRequest,
) (*petv1.PutPetResponse, error) {
	pet := &petv1.Pet{
		Name:    req.GetName(),
		PetType: req.PetType,
	}
	log.Printf("PutPet recieved a %v named %s\n", pet.GetPetType(), pet.GetName())
	return &petv1.PutPetResponse{Pet: pet}, nil
}
