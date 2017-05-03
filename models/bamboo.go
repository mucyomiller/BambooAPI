package models

import (
	"errors"

	"github.com/mucyomiller/BambooAPI/config"
	"gopkg.in/mgo.v2/bson"
)

// Menu struct
type Menu struct {
	ID          bson.ObjectId `json:"_id,omitempty"           bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty"          bson:"name,omitempty"`
	Description string        `json:"description,omitempty"   bson:"description,omitempty"`
	Price       float64       `json:"price,omitempty"         bson:"price,omitempty"`
	ImageURL    string        `json:"imageUrl,omitempty"      bson:"imageUrl,omitempty"`
}

//IsValid ...
func (m Menu) IsValid() bool {
	if l := len(m.Name); l > 3 {
		return true
	}

	return false
}

//Menus containing slices of Menu
type Menus []Menu

const col string = "menus"

//All get All Menus from DB
func All() (Menus, error) {
	db := config.DB{}
	ms := Menus{}

	s, err := db.DoDial()

	if err != nil {
		return ms, errors.New("there was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&ms)

	if err != nil {
		return ms, errors.New("there was an error trying to find the menus")
	}

	return ms, err
}

//GetByID func get Menu By Id
func GetByID(id string) (Menu, error) {
	db := config.DB{}
	m := Menu{}

	s, err := db.DoDial()

	if err != nil {
		return m, errors.New("there was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&m)

	if err != nil {
		return m, errors.New("there was an error trying to find the menus")
	}

	return m, err
}

//NewMenu create New Menu
func NewMenu(m Menu) (Menu, error) {
	db := config.DB{}
	m.ID = bson.NewObjectId()

	s, err := db.DoDial()

	if err != nil {
		return m, errors.New("there was an error trying to connect to the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Insert(m)

	if err != nil {
		return m, errors.New("there was an error trying to insert the doc to the DB")
	}

	return m, err
}

//DeleteMenu func to delete Menu
func DeleteMenu(id string) error {
	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("there was an error trying to connect with the DB")
	}

	idoi := bson.ObjectIdHex(id)

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.RemoveId(idoi)

	if err != nil {
		return errors.New("there was an error trying to remove the todo")
	}

	return err
}
