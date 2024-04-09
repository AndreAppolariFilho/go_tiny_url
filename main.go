package main

import (
	"github.com/AndreAppolariFilho/go_tiny_url/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){
	/*
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not found in the environment")
	}
	fmt.Println("Port: ",portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DB_URL is not found in the environment")
	}
	fmt.Println("DB_URL: ",dbURL)

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}
	*/
	//queries := database.New(conn)

	/*apiCfg := apiConfig{
		DB: queries,
	}*/
	/*
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server starting on port: %v", portString)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	*/
}