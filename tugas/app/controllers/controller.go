package controllers

import (
	"log"
	"net/http"
	"tugasmvc2/app/models"
	"tugasmvc2/app/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

	var account models.Account
	if err := c.Bind(&account); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
	pass, err := utils.HashGenerator(account.Password)
	if err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
	account.Password = pass
	flag, err := models.InsertNewAccount(account)
	if flag {
		utils.WrapAPISuccess(c, "success", http.StatusOK)
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetAccountMutasi(c *gin.Context) {
	idAccount := c.MustGet("account_number").(int)
	flag, err, trx, acc := models.GetAccountMutasiLast(idAccount)
	if err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if flag {
		utils.WrapAPIData(c, map[string]interface{}{
			"account":  acc,
			"mutation": trx,
		}, http.StatusOK, "success")
		return
	}
}

func GetAccount(c *gin.Context) {
	idAccount := c.MustGet("account_number").(int)
	flag, err, trx, acc := models.GetAccountDetail(idAccount)
	if err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if flag {
		utils.WrapAPIData(c, map[string]interface{}{
			"account":     acc,
			"transaction": trx,
		}, http.StatusOK, "success")
		return
	}
}

// func AddInterest (c *gin.Context){

// }

func Transfer(c *gin.Context) {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := models.Transfer(transaction)
	if flag {
		utils.WrapAPISuccess(c, "success", http.StatusOK)
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func Withdraw(c *gin.Context) {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := models.Withdraw(transaction)
	if flag {
		utils.WrapAPISuccess(c, "success", http.StatusOK)
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func Deposit(c *gin.Context) {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := models.Deposit(transaction)
	if flag {
		utils.WrapAPISuccess(c, "success", http.StatusOK)
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func Interest(c *gin.Context) {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := models.Interest(transaction)
	if flag {
		utils.WrapAPISuccess(c, "success", http.StatusOK)
		return
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
}

func Login(c *gin.Context) {
	var auth models.Auth
	if err := c.Bind(&auth); err != nil {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("LOGIN")
	flag, err, token := models.Login(auth)
	if flag {
		utils.WrapAPIData(c, map[string]interface{}{
			"token": token,
		}, http.StatusOK, "success")
	} else {
		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
	}
}
