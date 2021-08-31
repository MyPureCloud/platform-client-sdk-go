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

func (o *Leaderboarditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Leaderboarditem
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		
		Points *int `json:"points,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Rank: o.Rank,
		
		Points: o.Points,
		Alias:    (*Alias)(o),
	})
}

func (o *Leaderboarditem) UnmarshalJSON(b []byte) error {
	var LeaderboarditemMap map[string]interface{}
	err := json.Unmarshal(b, &LeaderboarditemMap)
	if err != nil {
		return err
	}
	
	if User, ok := LeaderboarditemMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Rank, ok := LeaderboarditemMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	
	if Points, ok := LeaderboarditemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Leaderboarditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
