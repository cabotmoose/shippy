package main

import (
	"context"
	
	pb "github.com/cabotmoose/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// MongoRepository implementation
type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}

	var vessel *pb.Vessel
	if err := repository.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create - Create new vessel
func (repository *VesselRepository) Create(vessel *pb.Vessel) error {
	return repository.collection.Insert(vessel)
}
