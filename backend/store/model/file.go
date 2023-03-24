package model

import "errors"

const (
	FileTypeImage = "image"
	FileTypeVideo = "video"
)

type File struct {
	ID   int    `db:"id" json:"-"`
	UID  string `db:"uid" json:"uid"`
	Path string `db:"path" json:"path"`
	Type string `db:"type" json:"type"`
}

func (f File) GetUID() string {
	return f.UID
}

func (f File) SetID(id int) {
	f.ID = id
}

func (f File) Verify() error {
	if f.Path == "" {
		return errors.New("missing path")
	} else if f.Type != FileTypeImage && f.Type != FileTypeVideo {
		return errors.New("missing or invalid type")
	}
	return nil
}
