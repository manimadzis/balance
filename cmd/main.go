package main

import (
	"balance/internal/balance/config"
	"balance/internal/balance/server"
	"balance/internal/balance/store/postgres"
	"database/sql"
	"flag"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("f", "", "Путь до конфигурационного файла")
	flag.Parse()

	var cfg *config.Config
	var err error
	if *configPath != "" {
		if cfg, err = config.ParseFile(*configPath); err != nil {
			log.Printf("Невозможно загрузить указанный конфиг: %v\n", err)
		}
	} else {
		cfg = config.DefaultConfig()
	}

	println(cfg.DBConnectURL)
	println(cfg.Address)

	db, err := sql.Open("postgres", cfg.DBConnectURL)
	if err != nil {
		log.Fatal(err)
	}
	store := postgres.New(db)

	serverConfig := &server.Config{
		Address: cfg.Address,
	}

	APIserver := server.New(serverConfig, store)
	if err := APIserver.Start(); err != nil {
		log.Fatal(err)
	}
}
