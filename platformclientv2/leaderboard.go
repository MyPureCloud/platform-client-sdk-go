package platformclientv2
import (
	"time"
	"encoding/json"
)

// Leaderboard
type Leaderboard struct { 
	// Division - The targeted division for this leaderboard
	Division *Division `json:"division,omitempty"`


	// Metric - The metric id if the leaderboard is about a specific metric
	Metric *Metric `json:"metric,omitempty"`


	// DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// Leaders - The list of leaders generated.
	Leaders *[]Leaderboarditem `json:"leaders,omitempty"`

}

// String returns a JSON representation of the model
func (o *Leaderboard) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
