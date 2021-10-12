package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT       = 0
	DBPAGESIZE = 0
	DSN        = ""
	DBDRIVER   = ""
	SECRETKEY  []byte
)

func Load() {
	// var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		log.Println(err)
		PORT = 9000
	}
	DBPAGESIZE, err = strconv.Atoi(os.Getenv("DB_PAGESIZE"))

	if err != nil {
		log.Println(err)
		DBPAGESIZE = 20
	}
	DBDRIVER = os.Getenv("DB_DRIVER")

	// DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&&parsTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),

		os.Getenv("DB_NAME"))

	SECRETKEY = []byte(os.Getenv("API_SECRET"))

}
