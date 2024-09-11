package db

import (
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supago"
)

var Supabase *supa.Client

func CreateClient(){
	err:=godotenv.Load(".env")

	// create client
	url := os.Getenv(("SUPABASE_URL"))
	key := os.Getenv(("SUPABASE_KEY"))
	Supabase = supa.CreateClient(url, key)

}