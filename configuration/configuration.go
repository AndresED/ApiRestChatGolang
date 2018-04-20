package configuration

// MYSQL => go get -u github.com/go-sql-driver/mysql
//GORM =>  go get -u github.com/jinzhu/gorm
import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Configuration Struct config bd
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

//GetConfiguration , Permite obtener los datos almacenados en el configuration.json para la conexion a la  bd
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

//GetConnection Devuelve la conexion a la bd
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	//data source name
	//  user:password@tcp(server:port)/database?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error en la conexión a la bd")
	}
	return db
}
