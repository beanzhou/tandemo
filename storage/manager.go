package storage

import (
	log "github.com/sirupsen/logrus"
)

type UserManager struct{}

func (m *UserManager) Insert(user User) (u User, err error) {
	err = db.Insert(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	u = user
	return
}

func (m *UserManager) Update(user User) (err error) {
	err = db.Update(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManager) Delete(userid int64) (err error) {
	user := User{Id: userid}
	err = db.Delete(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManager) Get(userid int64) (user User, err error) {
	user = User{Id: userid}
	err = db.Select(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

// TODO: Pagination
func (m *UserManager) GetUsers() (users []User, err error) {
	err = db.Model(&users).Order("id").Select()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManager) Count() (count int, err error) {
	count, err = db.Model(&User{}).Count()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

type RelationshipManager struct{}

func (m *RelationshipManager) Insert(rs Relationship) (r Relationship, err error) {
	err = db.Insert(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	r = rs
	return
}

func (m *RelationshipManager) Update(rs Relationship) (err error) {
	err = db.Update(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *RelationshipManager) Delete(rsId int64) (err error) {
	rs := Relationship{Id: rsId}
	err = db.Delete(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

// TODO: Pagination
func (m *RelationshipManager) GetRelationships(userid int64) (rs []Relationship, err error) {
	err = db.Model(&rs).Where("user_id = ?", userid).Order("id").Select()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *RelationshipManager) Count(userid int64) (count int, err error) {
	count, err = db.Model(&Relationship{}).Where("user_id = ?", userid).Count()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}
