package controller

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/SaKu2110/oauth_chain/wallet/model"
	"github.com/SaKu2110/oauth_chain/wallet/wallet"
)

type IsController struct {
	DB	*gorm.DB
	Wallet	*wallet.Wallet
}

func requestCheckSums(request model.Transaction) error{
	if request.SenderAddress == "" {
		return fmt.Errorf("SenderAddress value did not exist")
	}
	if request.ReceiverAddress == "" {
		return fmt.Errorf("ReceiverAddress value did not exist")
	}
	if request.Amount == 0.0 {
		return fmt.Errorf("Amount value is 0.0")
	}
	return nil
}

func (ctrl *IsController)CreateTransaction(context *gin.Context) {
	var transactionRequest model.Transaction
	var block []byte

	err := context.BindJSON(&transactionRequest)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}
	err = requestCheckSums(transactionRequest)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}
	block, err = ctrl.Wallet.CreateNewTransaction(transactionRequest)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	network := model.NetWork{
		Transaction: block,
	}
	context.JSON(http.StatusOK, network)
}