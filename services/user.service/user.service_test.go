package user_service_test

import (
	m "Golang-Mongodb/models"
	userService "Golang-Mongodb/services/user.service"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID:         oid,
		Name:       "Lis",
		Email:      "test@mail.com",
		CreatedAt:  time.Now(),
		UpdatedtAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("La creación del usuario a fallado: ", err)
		t.Fail()
	} else {
		t.Log("Usuario creado con éxito")
	}
}
func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("La consulta de usuarios a fallado: ", err)
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("La consulta no retornó datos: ", err)
		t.Fail()
	} else {
		t.Log("La consulta finalizó con éxito")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Lis Jardim",
		Email: "test_mongo@mail.com",
	}

	err := userService.Update(user, userId)

	if err != nil {
		t.Error("Error al actualizar al usuario: ", err)
		t.Fail()
	} else {
		t.Log("Usuario actualizado con éxito")
	}

}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)
	if err != nil {
		t.Error("Error al eliminar al usuario: ", err)
		t.Fail()
	} else {
		t.Log("Usuario eliminado con éxito")
	}

}
