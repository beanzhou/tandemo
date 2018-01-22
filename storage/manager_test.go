package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCRUD(t *testing.T) {
	userList := []User{
		{Id: 1, Name: "test1", Type: "user"},
		{Id: 2, Name: "test2", Type: "user"},
		{Id: 3, Name: "test3", Type: "user"},
	}

	userListAfterUpdate := []User{
		{Id: 1, Name: "test1", Type: "user"},
		{Id: 2, Name: "test22", Type: "user"},
		{Id: 3, Name: "test3", Type: "user"},
	}

	userListAfterDelete := []User{
		{Id: 2, Name: "test22", Type: "user"},
		{Id: 3, Name: "test3", Type: "user"},
	}

	userManger := &UserManger{}
	var err error
	// Insert
	for _, user := range userList {
		err = userManger.Insert(user)
		assert.NoError(t, err)
	}

	count, err := userManger.Count()
	assert.NoError(t, err)
	assert.Equal(t, 3, count)

	// Get Users
	users, err := userManger.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, userList, users)

	// Get
	user, err := userManger.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, User{Id: 1, Name: "test1", Type: "user"}, user)

	//Update
	err = userManger.Update(User{Id: 2, Name: "test22", Type: "user"})
	assert.NoError(t, err)
	users, err = userManger.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, userListAfterUpdate, users)

	// Delete
	err = userManger.Delete(1)
	assert.NoError(t, err)
	users, err = userManger.GetUsers()
	assert.Equal(t, userListAfterDelete, users)

	err = userManger.Delete(2)
	assert.NoError(t, err)
	err = userManger.Delete(3)
	assert.NoError(t, err)
}

func TestRelationshipCRUD(t *testing.T) {
	rsManager := RelationshipManger{}
	rsList := []Relationship{
		{Id: 1, UserId: 1, OtherId: 2, State: "liked", Type: "relationship"},
		{Id: 2, UserId: 1, OtherId: 3, State: "disiked", Type: "relationship"},
		{Id: 3, UserId: 1, OtherId: 4, State: "matched", Type: "relationship"},
	}
	rsListUpdate := []Relationship{
		{Id: 1, UserId: 1, OtherId: 2, State: "matched", Type: "relationship"},
		{Id: 2, UserId: 1, OtherId: 3, State: "disiked", Type: "relationship"},
		{Id: 3, UserId: 1, OtherId: 4, State: "matched", Type: "relationship"},
	}
	rsListDelete := []Relationship{
		{Id: 2, UserId: 1, OtherId: 3, State: "disiked", Type: "relationship"},
		{Id: 3, UserId: 1, OtherId: 4, State: "matched", Type: "relationship"},
	}

	// Insert
	var err error
	for _, rs := range rsList {
		err = rsManager.Insert(rs)
		assert.NoError(t, err)
	}

	count, err := rsManager.Count(1)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)

	// get list
	list, err := rsManager.GetRelationships(1)
	assert.NoError(t, err)
	assert.Equal(t, rsList, list)

	err = rsManager.Update(Relationship{Id: 1, UserId: 1, OtherId: 2, State: "matched", Type: "relationship"})
	assert.NoError(t, err)
	list, err = rsManager.GetRelationships(1)
	assert.NoError(t, err)
	assert.Equal(t, rsListUpdate, list)

	// Delete
	err = rsManager.Delete(1)
	assert.NoError(t, err)
	list, err = rsManager.GetRelationships(1)
	assert.NoError(t, err)
	assert.Equal(t, rsListDelete, list)

	err = rsManager.Delete(2)
	assert.NoError(t, err)
	err = rsManager.Delete(3)
	assert.NoError(t, err)

}
