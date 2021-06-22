package models

import (
	"github.com/uadmin/uadmin"
)

// Folder !
type Folder struct {
	uadmin.Model
	Name     string
	Parent   *Folder
	ParentID uint
}
