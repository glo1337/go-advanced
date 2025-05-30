package user

import "order-api/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindBySessionId(sessionId string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) UpdateSessionId(userId uint, sessionId string) (*User, error) {
	var user User

	if err := repo.Database.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}

	user.SessionId = sessionId
	if err := repo.Database.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) UpdateToken(userId uint, token string) (*User, error) {
	var user User

	if err := repo.Database.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}

	user.Token = token
	if err := repo.Database.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
