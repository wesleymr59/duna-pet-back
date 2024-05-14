package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	ENV_FILE = ".env"
)
const projectDirName = "duna-pet-back"

type Config struct {
	Connection string
	Host       string
	Port       int
	Username   string
	Password   string
	Name       string
	Charset    string
}

func loadEnv() {
	fmt.Println("taaki")
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.WithFields(log.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Problem loading .env file")

		os.Exit(-1)
	}
}

func GetConfig() *Config {
	loadEnv()

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &Config{
		Connection: "mysql",
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUsername,
		Password:   dbPassword,
		Name:       dbName,
		Charset:    "utf8",
	}
}
