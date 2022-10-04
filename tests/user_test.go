package tests

import (
	"github.com/stretchr/testify/assert"
	. "keyi/models"
	. "keyi/utils"
	"testing"
	"time"
)

const domain = "fudan.edu.cn"

func TestRegister(t *testing.T) {
	email := "TestRegister@" + domain
	password := "123456"
	code := "123456"
	_ = SetCache(email, code, time.Second*1)
	data := Map{"email": email, "password": password, "code": code}

	testAPIModel[User](t, "POST", "/api/users", 201, data)
}

func TestModifyPassword(t *testing.T) {
	var user User
	_ = DB.First(&user, 1)
	user.Password = "123456"
	user.Email = "TestModifyPassword@" + domain
	_ = DB.Save(&user)

	code := "123456"
	_ = SetCache(user.Email, code, time.Second*1)

	newPassword := "654321"
	data := Map{"password": user.Password, "new_password": newPassword}
	user = testAPIModel[User](t, "PUT", "/api/users/1", 200, data)
	_ = DB.First(&user, 1)
	assert.Equal(t, newPassword, user.Password)

	newPassword = "newPassword"
	data = Map{"code": code, "new_password": newPassword}
	testAPIModel[User](t, "PUT", "/api/users/1", 200, data)
	_ = DB.First(&user, 1)
	assert.Equal(t, newPassword, user.Password)
}
