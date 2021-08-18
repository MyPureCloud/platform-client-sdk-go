package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Leaderboarditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Leaderboarditem

	

	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		
		Points *int `json:"points,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Rank: u.Rank,
		
		Points: u.Points,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Leaderboarditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
