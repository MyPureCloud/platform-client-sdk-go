package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Performanceprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Performanceprofile

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MetricOrders *[]string `json:"metricOrders,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ReportingIntervals *[]Reportinginterval `json:"reportingIntervals,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		MaxLeaderboardRankSize *int `json:"maxLeaderboardRankSize,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		Description: u.Description,
		
		MetricOrders: u.MetricOrders,
		
		DateCreated: DateCreated,
		
		ReportingIntervals: u.ReportingIntervals,
		
		Active: u.Active,
		
		MaxLeaderboardRankSize: u.MaxLeaderboardRankSize,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Performanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
