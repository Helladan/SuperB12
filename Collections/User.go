package Collections

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

// User entity
type User struct {
	ID             bson.ObjectId         `json:"id" bson:"_id,omitempty"` // User ID
	Login          string                `json:"login"`                   // User login name
	Name           string                `json:"name"`                    // User full name
	Email          string                `json:"email"`                   // User email
	SignInDate     time.Time             `json:"sign_in_date"`            // User sign in date
	LastSignUpDate time.Time             `json:"last_sign_up_date"`       // User last sign up date
	Password       []byte                `json:"password"`                // User password (bcrypt hash)
	GlobalRank     string                `json:"global_rank"`             // User global rank
	GroupID        map[bson.ObjectId]int `json:"groupID"`                 // If so, group(s) who user belong to --- group ID => rank in group
}
