package platformclientv2
import (
	"encoding/json"
)

// Leaderboarditem
type Leaderboarditem struct { 
	// User - The user object for this leaderboard rank
	User *Userreference `json:"user,omitempty"`


	// Rank - The rank of the user
	Rank *int `json:"rank,omitempty"`


	// Points - The points collected by the user
	Points *int `json:"points,omitempty"`

}

// String returns a JSON representation of the model
func (o *Leaderboarditem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
