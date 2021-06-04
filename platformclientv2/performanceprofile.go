package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Performanceprofile
type Performanceprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - A name for this performance profile
	Name *string `json:"name,omitempty"`


	// Division - The division for this performance profile associate to
	Division *Division `json:"division,omitempty"`


	// Description - A description about this performance profile
	Description *string `json:"description,omitempty"`


	// MetricOrders - Order of the associated metrics. The list should contain valid ids for metrics
	MetricOrders *[]string `json:"metricOrders,omitempty"`


	// DateCreated - Creation date for this performance profile. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ReportingIntervals - The reporting interval periods for this performance profile
	ReportingIntervals *[]Reportinginterval `json:"reportingIntervals,omitempty"`


	// Active - The flag for active profiles
	Active *bool `json:"active,omitempty"`


	// MaxLeaderboardRankSize - The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
	MaxLeaderboardRankSize *int `json:"maxLeaderboardRankSize,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Performanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
