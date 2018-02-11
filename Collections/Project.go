package Collections

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Project entity
type Project struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"` // Project ID
	Name         string        `json:"name"`                    // Project name
	Description  []byte        `json:"description"`             // Project Description
	CreationDate time.Time     `json:"creation_date"`           // Project creation date
	UserID       bson.ObjectId `json:"user_id"`                 // User who create the project
	GroupID      bson.ObjectId `json:"group_id"`                // If so, group who create the project
}
