package Collections

import "github.com/globalsign/mgo/bson"

// Group of users entity
type Group struct {
	ID          bson.Binary `json:"id" bson:"_id,omitempty"` // Group ID
	Name        string      `json:"name"`                    // Groupe name
	Description []byte      `json:"description"`             // Group Description
}
