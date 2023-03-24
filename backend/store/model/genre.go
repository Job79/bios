package model

import "errors"

type Genre struct {
	ID   int    `db:"id" json:"-"`
	UID  string `db:"uid" json:"uid"`
	Name string `db:"name" json:"name"`
}

func (g *Genre) GetUID() string {
	return g.UID
}

func (g *Genre) SetID(id int) {
	g.ID = id
}

func (g *Genre) Verify() error {
	if g.Name == "" {
		return errors.New("missing name")
	}
	return nil
}
