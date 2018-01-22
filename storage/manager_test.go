package storage

import (
	"testing"
)

func TestUserCRUD(t *testing.T) {
	userManger := &UserManger{}
	userManger.Insert(User{Name: "test", Type: "user"})
}

func TestRelationshipCRUD(t *testing.T) {

}
