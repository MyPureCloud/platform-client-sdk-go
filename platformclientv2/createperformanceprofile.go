package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createperformanceprofile
type Createperformanceprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - A name for this performance profile
	Name *string `json:"name,omitempty"`


	// Division - The associated division for this Performance Profile
	Division *Writabledivision `json:"division,omitempty"`


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


	// MemberCount - The number of members in this performance profile
	MemberCount *int `json:"memberCount,omitempty"`


	// MaxLeaderboardRankSize - The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
	MaxLeaderboardRankSize *int `json:"maxLeaderboardRankSize,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Createperformanceprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createperformanceprofile
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MetricOrders *[]string `json:"metricOrders,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ReportingIntervals *[]Reportinginterval `json:"reportingIntervals,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
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
		
		MemberCount: o.MemberCount,
		
		MaxLeaderboardRankSize: o.MaxLeaderboardRankSize,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Createperformanceprofile) UnmarshalJSON(b []byte) error {
	var CreateperformanceprofileMap map[string]interface{}
	err := json.Unmarshal(b, &CreateperformanceprofileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateperformanceprofileMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CreateperformanceprofileMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := CreateperformanceprofileMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := CreateperformanceprofileMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if MetricOrders, ok := CreateperformanceprofileMap["metricOrders"].([]interface{}); ok {
		MetricOrdersString, _ := json.Marshal(MetricOrders)
		json.Unmarshal(MetricOrdersString, &o.MetricOrders)
	}
	
	if dateCreatedString, ok := CreateperformanceprofileMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ReportingIntervals, ok := CreateperformanceprofileMap["reportingIntervals"].([]interface{}); ok {
		ReportingIntervalsString, _ := json.Marshal(ReportingIntervals)
		json.Unmarshal(ReportingIntervalsString, &o.ReportingIntervals)
	}
	
	if Active, ok := CreateperformanceprofileMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if MemberCount, ok := CreateperformanceprofileMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if MaxLeaderboardRankSize, ok := CreateperformanceprofileMap["maxLeaderboardRankSize"].(float64); ok {
		MaxLeaderboardRankSizeInt := int(MaxLeaderboardRankSize)
		o.MaxLeaderboardRankSize = &MaxLeaderboardRankSizeInt
	}
	
	if SelfUri, ok := CreateperformanceprofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createperformanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
