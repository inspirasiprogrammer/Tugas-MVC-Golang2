package utils

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func HashGenerator(str string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedString), nil
}

func HashComparator(hashedByte []byte, byte []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedByte, byte)
	if err != nil {
		return err
	}
	return nil
}

func WrapAPIError(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"code":          code,
		"error_type":    http.StatusText(code),
		"error_details": message,
	})
}

func WrapAPISuccess(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"code":   code,
		"status": message,
	})
}
func WrapAPIData(c *gin.Context, data interface{}, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"code":   code,
		"status": message,
		"data":   data,
	})
}
