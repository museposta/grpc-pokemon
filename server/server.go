
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	pokemonpc "github.com/museposta/grpc-pokemon/pokemon"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	)
	var collection *mongo.Collection
	type server struct {
	  pokemonpc.PokemonServiceServer
	}
	type pokemonItem struct {
	 ID          primitive.ObjectID `bson:"_id,omitempty"`
	 Pid         string             `bson:"pid"`
	 Name        string             `bson:"name"`
	 Power       string             `bson:"power"`
	 Description string             `bson:"description"`
	}
	func getPokemonData(data *pokemonItem) *pokemonpc.Pokemon {
	 return &pokemonpc.Pokemon{
	  Id:          data.ID.Hex(),
	  Pid:         data.Pid,
	  Name:        data.Name,
	  Power:       data.Power,
	  Description: data.Description,
	 }
	}
	func (*server) CreatePokemon(ctx context.Context, req *pokemonpc.CreatePokemonRequest) (*pokemonpc.CreatePokemonResponse, error) {
	fmt.Println("Create Pokemon")
	pokemon := req.GetPokemon()
	data := pokemonItem{
	 Pid:         pokemon.GetPid(),
	 Name:        pokemon.GetName(),
	 Power:       pokemon.GetPower(),
	 Description: pokemon.GetDescription(),
	}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
	 return nil, status.Errorf(
	  codes.Internal,
	  fmt.Sprintf("Internal error: %v", err),
	 )
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	  if !ok {
	   return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Cannot convert to OID"),
	  )
	}
	return &pokemonpc.CreatePokemonResponse{
	  Pokemon: &pokemonpc.Pokemon{
	  Id:          oid.Hex(),
	  Pid:         pokemon.GetPid(),
	  Name:        pokemon.GetName(),
	  Power:       pokemon.GetPower(),
	  Description: pokemon.GetDescription(),
	  },
	}, nil
	}