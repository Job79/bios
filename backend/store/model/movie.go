package model

import (
	"errors"
	"time"
)

const (
	MovieStatusNow   = "now"
	MovieStatusSoon  = "soon"
	MovieStatusEnded = "ended"
)

type Movie struct {
	ID          int       `db:"id" json:"-"`
	UID         string    `db:"uid" json:"uid"`
	Name        string    `db:"name" json:"name"`
	Description *string   `db:"description" json:"description,omitempty"`
	TotalTime   int       `db:"total_time" json:"total_time,omitempty"`
	ReleaseDate time.Time `db:"release_date" json:"release_date,omitempty"`
	Status      string    `db:"status" json:"status,omitempty"`

	Genres          []*Genre `json:"genres,omitempty"`
	Classifications []*Movie `json:"classifications,omitempty"`
	Files           []*File  `json:"files,omitempty"`
}

func (m *Movie) GetUID() string {
	return m.UID
}

func (m *Movie) SetID(id int) {
	m.ID = id
}

func (m *Movie) Verify() error {
	if m.Name == "" {
		return errors.New("missing name")
	} else if m.TotalTime == 0 {
		return errors.New("missing total_time")
	} else if m.ReleaseDate.IsZero() {
		return errors.New("missing release_date")
	} else if m.Status != MovieStatusNow && m.Status != MovieStatusSoon && m.Status != MovieStatusEnded {
		return errors.New("missing or invalid status")
	}
	return nil
}
