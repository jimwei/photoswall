package respository

import (
	"errors"
	"log"
	common "photoswall/api/common"
	entity "photoswall/api/entity"
)

var (
	INSERTERROR = errors.New("Create New User Fail.")
)

type UserRespository struct{}

// get all users
func (this *UserRespository) GetAll() ([]*entity.User, error) {
	users := []*entity.User{}
	common.DB.Find(&users)
	return users, nil
}

// get the user via name
func (this *UserRespository) Get(name string) (*entity.User, error) {
	user := entity.User{}
	common.DB.First(&user)

	return &user, nil
}

// update the user
func (this *UserRespository) Update(user *entity.User) (bool, error) {
	originUser := entity.User{}
	originUser.ID = user.ID
	common.DB.First(&originUser)
	log.Print("the origin user is", originUser)
	//update property
	originUser.Name = user.Name
	originUser.Alias = user.Alias
	originUser.Password = user.Password

	//update
	common.DB.Save(&originUser)

	return true, nil
}

// insert a new user
func (this *UserRespository) Insert(user *entity.User) (*entity.User, error) {
	ret := common.DB.NewRecord(user)
	if ret {
		common.DB.Create(user)
	} else {
		return nil, INSERTERROR
	}
	return user, nil
}
