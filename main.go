package main

import (
	"fmt"
	"os"
	"strconv"

	"gormm/datastore"
	"gormm/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Host     string
	Port     int16
	DB       string
}

func ReadEnv() Config {
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("ERROR LOAD FILE", err)
	}

	res := Config{}
	res.Username = os.Getenv("User")
	res.DB = os.Getenv("Db")
	res.Host = os.Getenv("Host")
	intConv, _ := strconv.Atoi(os.Getenv("Port"))
	res.Port = int16(intConv)

	return res

}

func ConnectDB(configData Config) *gorm.DB {

	connString := fmt.Sprintf("%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configData.Username,
		configData.Host,
		configData.Port,
		configData.DB)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("Terjadi kesalahan koneksi database", err)
		return nil
	}
	return db
}

func main() {
	config := Config{"root", "localhost", 3306, "altagorm8db"}
	conn := ConnectDB(config)
	// conn.AutoMigrate(&Barang{})
	// fmt.Println(conn)
	// fmt.Println(conn.Error)

	userAcc := datastore.UserDB{Db: conn}
	barangAcc := datastore.BarangDB{Db: conn}

	// conn.AutoMigrate(&Transaksi{})
	// fmt.Println(conn)
	// fmt.Println(conn.Error)

	allUser, err := userAcc.GetAllDataUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allUser)

	_, err = barangAcc.InsertBarang(entities.Barang{Nama: "Sepatu"}) //dipanggil saat nambah lagi
	if err != nil {
		fmt.Println(err)
	}

	allBarang, err := barangAcc.GetAllDataBarang()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allBarang)
}
