package config

import(
	"os"
	"fmt"
	"log"
)

type dataBaseConfig struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

var c dataBaseConfig

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func checkElements(c dataBaseConfig) error {
	if c.User == "" {
		return fmt.Errorf("DB_USER value did not exist")
	}
	if c.Pass == "" {
		return fmt.Errorf("DB_PASS value did not exist")
	}
	if c.IP == "" {
		return fmt.Errorf("DB_IP value did not exist")
	}
	if c.Port == "" {
		return fmt.Errorf("DB_PORT value did not exist")
	}
	if c.Name == "" {
		return fmt.Errorf("DB_NAME value did not exist")
	}
	return nil
}

func GetConnectionToken() string {
	c = dataBaseConfig {
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		IP	: os.Getenv("DB_IP"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	err := checkElements(c)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}