package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongo "github.com/civilcoder55/go-lab/database"
	"github.com/civilcoder55/go-lab/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetAllGamesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)

	cursor, err := mongo.DB.Collection("games").Find(mongo.Ctx, bson.M{})

	if err != nil {
		encoder.Encode(map[string]string{"result": err.Error()})
		return
	}

	defer cursor.Close(mongo.Ctx)

	games := []models.Game{}
	for cursor.Next(mongo.Ctx) {
		var game models.Game
		err := cursor.Decode(&game)

		if err != nil {
			fmt.Println(err)
		} else {
			games = append(games, game)

		}
	}

	fmt.Println(games)
	encoder.Encode(games)
}

func GetOneGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	g := getGameById(params["id"])
	if g != nil {
		json.NewEncoder(w).Encode(g)
		return
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(map[string]string{"result": "not found"})
}

func getGameById(id string) *models.Game {
	var game models.Game

	_id, _ := bson.ObjectIDFromHex(id)
	err := mongo.DB.Collection("games").FindOne(mongo.Ctx, bson.M{"_id": _id}).Decode(&game)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &game
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"result": "empty body"})
		return
	}
	var game models.Game

	json.NewDecoder(r.Body).Decode(&game)

	if game.IsEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"result": "name couldn't be empty"})
		return
	}

	insertedGame, err := mongo.DB.Collection("games").InsertOne(mongo.Ctx, game)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"result": err.Error()})
	}

	json.NewEncoder(w).Encode(insertedGame)
}

func UpdateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	g := getGameById(params["id"])
	if g == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{"result": "not found"})
		return
	}

	if r.Body == nil {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(map[string]string{"result": "empty body"})
		return
	}
	var dto models.Game

	json.NewDecoder(r.Body).Decode(&dto)

	if dto.IsEmpty() {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(map[string]string{"result": "name couldn't be empty"})
		return
	}

	g.Name = dto.Name

	if dto.Year != 0 {
		g.Year = dto.Year
	}

	_, err := mongo.DB.Collection("games").UpdateOne(mongo.Ctx, bson.M{"_id": g.Id}, bson.D{{Key: "$set", Value: mongo.StructToBson(g)}})

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"result": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(g)
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	g := getGameById(params["id"])
	if g == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{"result": "not found"})
		return
	}

	_, err := mongo.DB.Collection("games").DeleteOne(mongo.Ctx, bson.M{"_id": g.Id})

	if err != nil {
		w.WriteHeader(500)

		json.NewEncoder(w).Encode(map[string]string{"result": err.Error()})
		return
	}

	w.WriteHeader(204)
}
