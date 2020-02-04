package main

import(
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/oauth_chain/wallet/config"
	"github.com/SaKu2110/oauth_chain/wallet/controller"
	"github.com/SaKu2110/oauth_chain/wallet/model"
	w "github.com/SaKu2110/oauth_chain/wallet/wallet"
)

func main(){
	db, err := initializeDataBase()
	if err != nil {
		log.Fatalf("failed initialize db. err=%s", err)
	}

	defer db.Close()

	var wallet *w.Wallet
	wallet, err = w.InitializeWallet()
	if err != nil {
		log.Fatalf("failed initialize db. err=%s", err)
	}

	ctrl := initializeController(db, wallet)
	router := setupRouter(ctrl)
	err = router.Run(":9005")
	if err != nil {
		log.Fatalf("failed launch router. err=%s", err)
	}
}

func initializeDataBase() (*gorm.DB, error){
	var db *gorm.DB
	var err error
	var count time.Duration
	token := config.GetConnectionToken()

	count = 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("faild mysql connection")
		}
		db, err = gorm.Open("mysql", token)
		if err == nil {
			db.AutoMigrate(&model.Hash{})
			return db, nil
		}
		time.Sleep(3 * time.Second)

		count++
	}

	return nil, err
}

func initializeController(db *gorm.DB, wallet *w.Wallet) (controller.IsController){
	return controller.IsController{
		DB: db,
		Wallet: wallet,
	}
}

func setupRouter(ctrl controller.IsController) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "This is Card-Info API"})
	})
	router.POST("/transaction", ctrl.CreateTransaction)
	return router
}