package Collections

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Comments under issues entity
type IssueComment struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"` // ID
	Content       []byte        `json:"content"`                 // Content
	PostDate      time.Time     `json:"post_date"`               // Post date
	ModDate       time.Time     `json:"mod_date"`                // Modify date
	UserID        bson.ObjectId `json:"user_id"`                 // User ID who post the comment
	ParentComment bson.ObjectId `json:"parent_comment"`          // If so, parent comment ID
}
