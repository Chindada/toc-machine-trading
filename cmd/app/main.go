package main

import (
	"os"
	"path/filepath"

	"tmt/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(filepath.Dir(ex), ".env"))
	if err != nil {
		panic(err)
	}

	app.InitDB()
	app.RunApp()
}
