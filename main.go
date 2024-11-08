package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	quiz "main/quizzes"
	"main/routes"

	"github.com/joho/godotenv"
)



type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}



func main(){

	r := gin.Default()

	dbConfig, err := LoadDBConfig()

 if err != nil {
   log.Fatal(err)
 }

	dsn := dbConfig.GetDSN()


	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("Connected to the database")
	queries := quiz.New(db)


	

	r.POST("/api/questions", routes.CreateQuestion(queries))
	r.PUT("/api/questions/:id", routes.UpdateQuestion(queries))
	r.GET("/api/questions/:id", routes.GetRandomQuestions(queries))
    r.GET("/api/questionsfromquiz/:id", routes.GetQuizWithQuestions(queries))



	r.POST("/api/quiz", routes.CreateQuiz(queries))
	r.GET("/api/quiz/:id", routes.GetQuiz(queries))
	r.GET("/api/quiz", routes.ListQuizzes(queries))
   
	r.Run() 

}




func LoadDBConfig() (*DBConfig, error) {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }

    config := &DBConfig{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("SSL_MODE"),
    }

    return config, nil
}

func (config *DBConfig) GetDSN() string {
    return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
        config.User,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
        config.SSLMode,
    )
}