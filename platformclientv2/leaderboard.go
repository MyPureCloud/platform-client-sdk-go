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

func (o *Leaderboard) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Leaderboard
	
	DateStartWorkday := new(string)
	if o.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(o.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if o.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(o.DateEndWorkday, "%Y-%m-%d")
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
		Division: o.Division,
		
		Metric: o.Metric,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		Leaders: o.Leaders,
		
		UserRank: o.UserRank,
		Alias:    (*Alias)(o),
	})
}

func (o *Leaderboard) UnmarshalJSON(b []byte) error {
	var LeaderboardMap map[string]interface{}
	err := json.Unmarshal(b, &LeaderboardMap)
	if err != nil {
		return err
	}
	
	if Division, ok := LeaderboardMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Metric, ok := LeaderboardMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if dateStartWorkdayString, ok := LeaderboardMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := LeaderboardMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if Leaders, ok := LeaderboardMap["leaders"].([]interface{}); ok {
		LeadersString, _ := json.Marshal(Leaders)
		json.Unmarshal(LeadersString, &o.Leaders)
	}
	
	if UserRank, ok := LeaderboardMap["userRank"].(map[string]interface{}); ok {
		UserRankString, _ := json.Marshal(UserRank)
		json.Unmarshal(UserRankString, &o.UserRank)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Leaderboard) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
