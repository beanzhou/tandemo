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

	userManager := &UserManager{}
	var err error
	// Insert
	for _, user := range userList {
		_, err = userManager.Insert(user)
		assert.NoError(t, err)
	}

	count, err := userManager.Count()
	assert.NoError(t, err)
	assert.Equal(t, 3, count)

	// Get Users
	users, err := userManager.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, userList, users)

	// Get
	user, err := userManager.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, User{Id: 1, Name: "test1", Type: "user"}, user)

	//Update
	err = userManager.Update(User{Id: 2, Name: "test22", Type: "user"})
	assert.NoError(t, err)
	users, err = userManager.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, userListAfterUpdate, users)

	// Delete
	err = userManager.Delete(1)
	assert.NoError(t, err)
	users, err = userManager.GetUsers()
	assert.Equal(t, userListAfterDelete, users)

	err = userManager.Delete(2)
	assert.NoError(t, err)
	err = userManager.Delete(3)
	assert.NoError(t, err)
}

func TestRelationshipCRUD(t *testing.T) {
	rsManager := RelationshipManager{}
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
		_, err = rsManager.Insert(rs)
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
