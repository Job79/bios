package model

import "errors"

type Classification struct {
	ID          int     `db:"id" json:"-"`
	UID         string  `db:"uid" json:"uid"`
	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description,omitempty"`
}

func (c *Classification) GetUID() string {
	return c.UID
}

func (c *Classification) SetID(id int) {
	c.ID = id
}

func (c *Classification) Verify() error {
	if c.Name == "" {
		return errors.New("missing name")
	}
	return nil
}
