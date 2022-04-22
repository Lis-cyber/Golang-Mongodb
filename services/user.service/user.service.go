package user_service

import (
	m "Golang-Mongodb/models"
	user_repository "Golang-Mongodb/repositories/user.repositories"
)

func Create(user m.User) error {
	err := user_repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
func Read() (m.Users, error) {
	users, err := user_repository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func Update(user m.User, userId string) error {
	err := user_repository.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}
func Delete(userId string) error {
	err := user_repository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
