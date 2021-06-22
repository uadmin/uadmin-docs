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
