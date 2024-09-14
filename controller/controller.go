package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/backend/model"
	"github.com/redis/go-redis/v9"

	// "github.com/redis/go-redis/v9"
	"github.com/backend/config"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://punamkumari399:R5ophqEn67VgUIiF@cluster0.5ykvx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const (
	dbName = "shuffle-Cards"
	colNameUsers = "users"
	colNameGames = "games"
)

var usersCollection  *mongo.Collection
var gamesCollection  *mongo.Collection


type Response struct{
	Message string `json:"message"`
	User model.User `json:"user"`
}


func init(){
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection successfull")
	usersCollection = client.Database(dbName).Collection(colNameUsers)
	gamesCollection = client.Database(dbName).Collection(colNameGames)
	fmt.Println("collection instance is ready")
}

// mongo helpers

// insert 1 user
func insertOneUser(user model.User){
	inserted, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("inserted one  user with id: ", inserted.InsertedID)
}

func getOneUser(name string) (*model.User, error)  {
	var user model.User
	err := usersCollection.FindOne(context.Background(), bson.M{"name": name}).Decode(&user)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("User is: ", user)
	return &user, nil
}


func getAllUsers() []primitive.M {
	cur, err := usersCollection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var users []primitive.M
	for cur.Next(context.Background()){
		var user bson.M
		err := cur.Decode(&user)
		if err !=nil{
			log.Fatal(err)
		}
		users = append(users, user)

	}

	defer cur.Close(context.Background())
	return users
}

func insertOneGame(game model.Game){
	ctx := context.Background()
	inserted, err := gamesCollection.InsertOne(context.Background(), game)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("game id: ", game.ID, "and game point: ", game.Point)

	if config.Rdb !=nil{
		fmt.Println("rdb is not null")
	}else if config.Rdb==nil{
		fmt.Println("rdb is null")
	}
	gamedetails :=map[string]interface{}{
		"_id" : primitive.NewObjectID(),
		"point":strconv.Itoa(game.Point),
	}
	jsonData, err := json.Marshal(gamedetails)
	if err != nil{
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	err = config.Rdb.HSet(ctx, "game:"+game.UserID, jsonData,0).Err()
	 if err != nil {
        // http.Error(w, "Error saving user to Redis", http.StatusInternalServerError)
        log.Println("Error saving to Redis:", err)
        return
    }

	fmt.Println("inserted One game with id: ", inserted.InsertedID)
}

func getOneGame(user_id string) (*model.Game, error)  {

	fmt.Println("user id in get one game: ", user_id)
	ctx := context.Background()
	id, _ := primitive.ObjectIDFromHex(user_id)
	var game model.Game

	data := gamesCollection.FindOne(context.Background(), bson.M{"_id": id})
	err := data.Decode(&game)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("data in get one game: ", game)

	val, err := config.Rdb.HGetAll(ctx, "game:"+game.UserID).Result()
	if err == redis.Nil{
		fmt.Println("game does not exist as key in redis")
	}else if err != nil{
		log.Fatalf("Error getting value from Redis : %v", err)
	} else {
		fmt.Println("val in get one game: ", val)
	}

	

	var retrievalGame model.Game
	for key := range val{
		err  := json.Unmarshal([]byte(key), &retrievalGame)
		if err != nil{
			return nil, fmt.Errorf("error unmarshalling Json %v", err)
		}
	}

	fmt.Println("the retrieval gaem : ", retrievalGame)





	// err = json.Unmarshal([]byte(val), &retrievalGame)
	// // err = json.Unmarshal([]byte(val["_id"]), &retrievalGame.ID)
	// // // err = json.Unmarshal([]byte(val["user_id"]), &retrievalGame.UserID)
	// // err = json.Unmarshal([]byte(val["point"]), &retrievalGame.Point)
	// if err != nil{
	// 	log.Fatalf("Error unmarshalling JSON: %s", err)
	// }

	// if retrievalGame.ID.Hex() != "" {
    //         fmt.Println("Game found in Redis:", retrievalGame)
    //         return &retrievalGame, nil
    //     }
    

    // If game not found in Redis, return the game from MongoDB
    fmt.Println("Game found in redis:", retrievalGame)
    return &retrievalGame, nil
}


func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user) // used to decode the req body received from an Http request into a go struct
	if err != nil{
		log.Fatal(err)
	}

	// preparing the response
	response := Response{
		Message: "User created successfully",
		User: user,
	}

	// encoding the response as json and sending it back to the client
	w.WriteHeader(http.StatusCreated)
	
	// insertOneUser(user) Insert the user into the database
	insertOneUser(user)
	err = json.NewEncoder(w).Encode(user) // encode the user as json and send it back to the client
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("user inserted info: ", response)
}

func GetOneUser(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "GET")
	params := mux.Vars(r)
	name := params["name"]
	user, err := getOneUser(name)
	if err != nil{
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)

	fmt.Println("user info: ", user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allUsers := getAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}



func CreateGame(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")
	// game := &model.Game{}
	var game model.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	insertOneGame(game)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func GetOneGame(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Content-Allow-Methods", "GET")
	
	params := mux.Vars(r)
    gameID, ok := params["id"]
    if !ok {
        http.Error(w, "Game ID is required", http.StatusBadRequest)
        return
    }

    game, err := getOneGame(gameID) // Assuming this function retrieves the game
    if err != nil {
        http.Error(w, "Game not found", http.StatusNotFound)
        return
    }

    if err := json.NewEncoder(w).Encode(game); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}