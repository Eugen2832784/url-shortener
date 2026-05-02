package main

import (
	"fmt"
	"log/slog"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
)

func main() {
	// TODO: implement config : cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)
	//TODO: init logger: slog

	//TODO:init storage: sqlite

	//TODO:init router: chi, "chi render"

	//TODO: run server:
}

func setupLogger(env string) {
	var log *slog.Logger
	switch env {
	case "local":

	}
}
