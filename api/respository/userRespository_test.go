package respository

import (
	entity "photoswall/api/entity"
	"testing"
)

// test get all func
func TestUserRespository_GetAll(t *testing.T) {
	userRes := new(UserRespository)

	users, err := userRes.GetAll()
	if err != nil {
		t.Error("get user fail.", err.Error())
	} else {
		t.Logf("the users is %v", users)
	}

}

// test insert a user func
func TestUserRespository_Insert(t *testing.T) {
	userRes := new(UserRespository)
	user := &entity.User{}
	user.Name = "jimwei"
	user.Alias = "jim"
	user.Password = "xA123456"
	user, err := userRes.Insert(user)
	if err != nil {
		t.Error("insert user fail.", err.Error())
	} else {
		t.Log("the insert user is", user)
	}
}

// test update user func
func TestUserRespository_Update(t *testing.T) {
	userRes := new(UserRespository)
	user := &entity.User{}
	user.ID = 2
	user.Name = "jimwei"
	user.Alias = "jim1"
	user.Password = "xA123456"
	ret, err := userRes.Update(user)
	if err != nil || !ret {
		t.Error("update user fail.", err.Error())
	} else {
		t.Log("update user success")
	}
}

func TestUserRespository_Get(t *testing.T) {
	userRes := new(UserRespository)

	user, err := userRes.Get("jimwei")
	if err != nil {
		t.Error("get user fail.", err.Error())
	} else {
		t.Log("get user successfully.", user)

	}
}
