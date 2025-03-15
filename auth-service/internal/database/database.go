package database

import (
	"auth-service/internal/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Service represents a service that interacts with a database.
type Service interface {
	Health() map[string]string
	Close() error
	GetUser(username string) (*model.User, error)
	Signup(userJson model.SignUp) (string,error)
	GetDeliveryDriver (username string) (*model.AuthDelivery,error)
}

type service struct {
	db *sql.DB
	mongoClient    *mongo.Client
	mongoColl  *mongo.Collection
	mongoCollDD *mongo.Collection
}

var (
	database   = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	schema     = os.Getenv("BLUEPRINT_DB_SCHEMA")
	mongoURI  = os.Getenv("MONGO_URI")  
	mongoDB   = os.Getenv("MONGO_DB")   
	mongoColl = os.Getenv("MONGO_COLL") 
	mongoCollDA = os.Getenv("MONGO_COLLDA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	clientOptions := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}


	mongoCollection := mongoClient.Database(mongoDB).Collection(mongoColl)
	mongoCollectionDeliveryDrivers := mongoClient.Database(mongoDB).Collection(mongoCollDA)

	dbInstance = &service{
		db: db,
		mongoClient:    mongoClient,
		mongoColl:  mongoCollection,
		mongoCollDD: mongoCollectionDeliveryDrivers,
	}
	return dbInstance
}

func (s *service) GetUser(username string) (*model.User, error) {
	var user model.User
	query := "SELECT username, password, role FROM users WHERE username = $1"
	err := s.db.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *service) GetDeliveryDriver(username string) (*model.AuthDelivery,error){
	filter := bson.M{"username": username} // You can modify the filter as needed
	var user *model.AuthDelivery
	err := s.mongoCollDD.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal("Error finding users:", err)
		return nil,err
	}
	return user,nil
}

func (s *service) Signup(userJson model.SignUp) (string,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3)"
	tmp,err := s.db.Exec(query,userJson.Username,userJson.Password,"user")
	if err != nil{
		log.Println(err)
	}
	log.Println(tmp)
	inserted,err := s.mongoColl.InsertOne(ctx,userJson);

	if err != nil{
		log.Fatal(err)
	}
	return fmt.Sprintf("Inserted 1 user with id %v", inserted.InsertedID),nil
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
