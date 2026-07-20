package repo

import (
	"time"

	"github.com/murashi19/koda-b8-backend1/internal/models"
)

type UserRepo struct {
	data *[]models.User
}

func NewUserRepo(data *[]models.User) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) Create(data *models.CreateUserRequest) *models.User {
	id := int64(len(*r.data) + 1)

	newUser := models.User{
		ID:        id,
		Email:     data.Email,
		Password:  data.Password,
		Username:  data.Username,
		Phone:     data.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	*r.data = append(*r.data, newUser)
	return &newUser
}

func (r *UserRepo) FindByEmail(email string) *models.User {
	for i := range *r.data {
		if (*r.data)[i].Email == email {
			return &(*r.data)[i]
		}
	}

	return nil
}

func (r *UserRepo) GetAllUsers() []models.User {
	return *r.data
}
