package helper_test

import (
	"final-project/config"
	"final-project/dto"
	"final-project/helper"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	test := dto.UserCreateResponse{
		ID:       1,
		Username: "test",
		Email:    "test@test.com",
	}
	config := config.Config()
	token, _ := helper.GenerateToken(test, config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	fmt.Println(token)
	assert.Equal(t, reflect.TypeOf(token), "string")
}
