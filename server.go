package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shirch/graphql/graph"
	"github.com/shirch/graphql/graph/generated"
	"github.com/shirch/graphql/graph/model"
	"github.com/shirch/graphql/internal/auth"

	"github.com/go-chi/chi"
)

var db *gorm.DB

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware(db))

	initDB()
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB() {
	var err error
	dataSourceName := "root:dbpass@tcp(localhost)/dbname"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	db.AutoMigrate(&model.Link{}, &model.User{})
}
