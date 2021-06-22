<<<<<<< HEAD
package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

// Todo Model !
type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	Category    Category
	CategoryID  uint
	Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
	FriendID    uint
	Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
	ItemID      uint
	TargetDate  time.Time
	Progress    int `uadmin:"progress_bar"`
}
=======
package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

// Todo Model !
type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	Category    Category
	CategoryID  uint
	Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
	FriendID    uint
	Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
	ItemID      uint
	TargetDate  time.Time
	Progress    int `uadmin:"progress_bar"`
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
