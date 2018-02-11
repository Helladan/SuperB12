package Collections

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Issues report entity
type Issue struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"` // Issue ID
	Title       string        `json:"title"`                   // Issue Title
	Description []byte        `json:"description"`             // Issue Description
	PostDate    time.Time     `json:"post_date"`               // Issue post date
	Status      int           `json:"statut"`                  // Issue status
	Priority    int           `json:"priority"`                // Issue priority
	UserID      bson.ObjectId `json:"userID"`                  // User who post the issue
	ProjectID   bson.ObjectId `json:"project_id"`              // Project that issue belong to
}
