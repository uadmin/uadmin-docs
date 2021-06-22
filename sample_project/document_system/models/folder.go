<<<<<<< HEAD
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
=======
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
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
