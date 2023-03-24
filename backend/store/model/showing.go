package model

import (
	"errors"
	"time"
)

type Showing struct {
	ID        int       `db:"id" json:"-"`
	UID       string    `db:"uid" json:"uid"`
	MovieID   int       `db:"movie_id" json:"-"`
	RoomID    int       `db:"room_id" json:"-"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`

	Room  *Room  `json:"room,omitempty"`
	Movie *Movie `json:"movie,omitempty"`
}

func (s *Showing) GetUID() string {
	return s.UID
}

func (s *Showing) SetID(id int) {
	s.ID = id
}

func (s *Showing) Verify() error {
	if s.StartTime.IsZero() {
		return errors.New("missing start_time")
	} else if s.StartTime.After(s.EndTime) {
		return errors.New("invalid start_time")
	} else if s.Room == nil || s.Room.UID == "" {
		return errors.New("missing room")
	} else if s.Movie == nil || s.Movie.UID == "" {
		return errors.New("missing movie")
	}
	return nil
}
