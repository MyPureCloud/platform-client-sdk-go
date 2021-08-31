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

func (o *Performanceprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Performanceprofile
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		MetricOrders: o.MetricOrders,
		
		DateCreated: DateCreated,
		
		ReportingIntervals: o.ReportingIntervals,
		
		Active: o.Active,
		
		MaxLeaderboardRankSize: o.MaxLeaderboardRankSize,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Performanceprofile) UnmarshalJSON(b []byte) error {
	var PerformanceprofileMap map[string]interface{}
	err := json.Unmarshal(b, &PerformanceprofileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PerformanceprofileMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PerformanceprofileMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := PerformanceprofileMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := PerformanceprofileMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if MetricOrders, ok := PerformanceprofileMap["metricOrders"].([]interface{}); ok {
		MetricOrdersString, _ := json.Marshal(MetricOrders)
		json.Unmarshal(MetricOrdersString, &o.MetricOrders)
	}
	
	if dateCreatedString, ok := PerformanceprofileMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ReportingIntervals, ok := PerformanceprofileMap["reportingIntervals"].([]interface{}); ok {
		ReportingIntervalsString, _ := json.Marshal(ReportingIntervals)
		json.Unmarshal(ReportingIntervalsString, &o.ReportingIntervals)
	}
	
	if Active, ok := PerformanceprofileMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if MaxLeaderboardRankSize, ok := PerformanceprofileMap["maxLeaderboardRankSize"].(float64); ok {
		MaxLeaderboardRankSizeInt := int(MaxLeaderboardRankSize)
		o.MaxLeaderboardRankSize = &MaxLeaderboardRankSizeInt
	}
	
	if SelfUri, ok := PerformanceprofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Performanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
