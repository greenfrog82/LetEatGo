package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	_ "os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "qwer1234"
	dbname   = "leteatgo"
)
const (
	secret_key = "CDNetworks2018"
)

type User struct {
	Username string `json: username`
	Password string `json: password`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//Migrate the schema
	// db.DropTable(&AuthUser{})
	db.AutoMigrate(&AuthUser{})

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/signup", signup(db))
	router.POST("/login", login(db))

	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}

type (
	AuthUser struct {
		gorm.Model        // Add fields `ID`, `CreateAt`, `UpdateAt`, `DeletedAd`
		Username   string `gorm:"unique_index;type:VARCHAR(100);NOT NULL;DEFAULT: ''"`
		Password   string `gorm:"type:VARCHAR(100);NOT NULL;DEFAULT: ''"`
		Useflag    bool   `gorm:"type:BOOLEAN;NOT NULL;DEFAULT: true"`
	}
)

func signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		h := sha256.New()
		h.Write([]byte(u.Password))
		sha256_passwd := hex.EncodeToString(h.Sum(nil))

		row := db.Create(&AuthUser{Username: u.Username, Password: sha256_passwd})
		if row.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Duplicated Username."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		h := sha256.New()
		h.Write([]byte(u.Password))
		sha256_passwd := hex.EncodeToString(h.Sum(nil))

		var user AuthUser
		row := db.Where("username = ? AND password = ? AND useflag = true", u.Username, sha256_passwd).Find(&user)
		if row.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "Auth User does not exist."})
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		expire := time.Now().Add(time.Hour * 72).Unix() // 3 days

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["user_id"] = user.ID
		claims["exp"] = expire

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(secret_key))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "make token failed"})
			return
		}
		c.JSON(http.StatusOK, map[string]string{
			"token":  t,
			"expire": strconv.FormatInt(expire, 10),
		})
	}
}
