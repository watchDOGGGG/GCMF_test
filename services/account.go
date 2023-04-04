package service

import (
	helper "gcmf-services/helpers"
	"gcmf-services/model"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func Verifyuseraccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		var accounts []model.VerifyAccountModel
		if err := c.BindJSON(&accounts); err != nil {
			println("invalide request body")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		// Create a wait group to wait for all Goroutines to finish
		var wg sync.WaitGroup

		// Create a channel to receive results from Goroutines
		resultChan := make(chan map[string]interface{})

		// Start a Goroutine for each account
		for _, account := range accounts {
			wg.Add(1)
			go func(account model.VerifyAccountModel) {
				defer wg.Done()

				// Convert the account variable to a struct
				acc := struct {
					AccountNumber string
					BankCode      string
				}{
					AccountNumber: account.AccountNumber,
					BankCode:      account.BankCode,
				}

				// Make API request to get account details
				details, err := helper.GetNameEnquiry(os.Getenv("AUTHTOKEN"), acc.AccountNumber, acc.BankCode)
				if err != nil {
					resultChan <- map[string]interface{}{
						"accountNumber": acc.AccountNumber,
						"bankCode":      acc.BankCode,
						"bankName":      account.BankName,
						"accountName":   account.AccountName,
						"amount":        account.Amount,
						"accountType":   account.AccountType,
						"status":        "failed",
						"error":         err.Error(),
					}
					return
				}

				resultChan <- map[string]interface{}{
					"accountNumber": acc.AccountNumber,
					"bankCode":      acc.BankCode,
					"bankName":      account.BankName,
					"accountName":   account.AccountName,
					"amount":        account.Amount,
					"accountType":   account.AccountType,
					"status":        "success",
					"data":          details,
				}
			}(account)
		}

		// Close the channel after all Goroutines have finished
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Collect results from channel and send response
		results := make([]map[string]interface{}, 0)
		for result := range resultChan {
			results = append(results, result)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Accounts verification completed",
			"data":    results,
		})

	}
}

func TestApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome"})
	}

}
