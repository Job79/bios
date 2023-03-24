package model

import "errors"

type Room struct {
	ID   int    `db:"id" json:"-"`
	UID  string `db:"uid" json:"uid"`
	Code string `db:"code" json:"code"`
}

func (r *Room) GetUID() string {
	return r.UID
}

func (r *Room) SetID(id int) {
	r.ID = id
}

func (r *Room) Verify() error {
	if r.Code == "" {
		return errors.New("missing code")
	}
	return nil
}
