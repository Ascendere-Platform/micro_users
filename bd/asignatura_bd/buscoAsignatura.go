package asignaturabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func BuscoAsignatura(id string) (models.Asignatura, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("asignatura")

	objID,_ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado models.Asignatura

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	
	if err != nil {
		return resultado, err
	}
	return resultado, err
}