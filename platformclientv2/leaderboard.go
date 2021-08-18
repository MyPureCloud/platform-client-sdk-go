package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Leaderboard
type Leaderboard struct { 
	// Division - The targeted division for this leaderboard
	Division *Division `json:"division,omitempty"`


	// Metric - The metric id if the leaderboard is about a specific metric
	Metric *Addressableentityref `json:"metric,omitempty"`


	// DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// Leaders - The list of leaders generated.
	Leaders *[]Leaderboarditem `json:"leaders,omitempty"`


	// UserRank - The requesting user's rank
	UserRank *Leaderboarditem `json:"userRank,omitempty"`

}

func (u *Leaderboard) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Leaderboard

	
	DateStartWorkday := new(string)
	if u.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(u.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if u.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(u.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	

	return json.Marshal(&struct { 
		Division *Division `json:"division,omitempty"`
		
		Metric *Addressableentityref `json:"metric,omitempty"`
		
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		Leaders *[]Leaderboarditem `json:"leaders,omitempty"`
		
		UserRank *Leaderboarditem `json:"userRank,omitempty"`
		*Alias
	}{ 
		Division: u.Division,
		
		Metric: u.Metric,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		Leaders: u.Leaders,
		
		UserRank: u.UserRank,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Leaderboard) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
