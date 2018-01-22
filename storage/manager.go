package storage

import (
	log "github.com/sirupsen/logrus"
)

type UserManger struct{}

func (m *UserManger) Insert(user User) (err error) {
	err = db.Insert(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManger) Update(user User) (err error) {
	err = db.Insert(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManger) Delete(userid int64) (err error) {
	user := User{Id: userid}
	err = db.Delete(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManger) GetOne(userid int64) (user User, err error) {
	user = User{Id: userid}
	err = db.Select(&user)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

// TODO: Pagination
func (m *UserManger) GetList(userid int64) (users []User, err error) {
	err = db.Model(&users).Order("id").Select()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *UserManger) Count() (count int, err error) {
	count, err = db.Model(&User{}).Count()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

type RelationshipManger struct{}

func (m *RelationshipManger) Insert(rs Relationship) (err error) {
	err = db.Insert(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *RelationshipManger) Update(rs Relationship) (err error) {
	err = db.Insert(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *RelationshipManger) Delete(rsId int64) (err error) {
	rs := Relationship{Id: rsId}
	err = db.Delete(&rs)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

// TODO: Pagination
func (m *RelationshipManger) GetRelationships(userid int64) (rs []Relationship, err error) {
	err = db.Model(&rs).Where("userid = ?", userid).Select()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (m *RelationshipManger) Count(userid int64) (count int, err error) {
	count, err = db.Model(&Relationship{}).Where("userid = ?", userid).Count()
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}
