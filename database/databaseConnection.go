package database

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func Connection() (*sql.DB, error) {
	database_name := viperEnvVariable("DATABASE_NAME")
	database_username := viperEnvVariable("DATABASE_USERNAME")
	database_password := viperEnvVariable("DATABASE_PASSWORD")

	stringConnectionDataBase := database_username + ":" + database_password + "@/" + database_name + "?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", stringConnectionDataBase)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err

	}
	return db, nil
}
