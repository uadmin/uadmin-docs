<<<<<<< HEAD
package models

import (
	"github.com/uadmin/uadmin"
)

// Nationality Field !
type Nationality int

// Chinese !
func (Nationality) Chinese() Nationality {
	return 1
}

// Filipino !
func (Nationality) Filipino() Nationality {
	return 2
}

// Others !
func (Nationality) Others() Nationality {
	return 3
}

// Friend Model !
type Friend struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Email       string `uadmin:"email"`
	Password    string `uadmin:"password;list_exclude"`
	Nationality Nationality
	Invite      string `uadmin:"link"`
}

// Save !
func (f *Friend) Save() {
	f.Invite = "https://www.google.com/"
	uadmin.Save(f)
}
=======
package models

import (
	"github.com/uadmin/uadmin"
)

// Nationality Field !
type Nationality int

// Chinese !
func (Nationality) Chinese() Nationality {
	return 1
}

// Filipino !
func (Nationality) Filipino() Nationality {
	return 2
}

// Others !
func (Nationality) Others() Nationality {
	return 3
}

// Friend Model !
type Friend struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Email       string `uadmin:"email"`
	Password    string `uadmin:"password;list_exclude"`
	Nationality Nationality
	Invite      string `uadmin:"link"`
}

// Save !
func (f *Friend) Save() {
	f.Invite = "https://www.google.com/"
	uadmin.Save(f)
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
