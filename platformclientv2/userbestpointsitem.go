package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userbestpointsitem
type Userbestpointsitem struct { 
	// GranularityType - Best points aggregation interval granularity
	GranularityType *string `json:"granularityType,omitempty"`


	// Points - Gamification points
	Points *int `json:"points,omitempty"`


	// DateStartWorkday - Start workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// Rank - The rank of this user
	Rank *int `json:"rank,omitempty"`

}

func (o *Userbestpointsitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userbestpointsitem
	
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
		GranularityType *string `json:"granularityType,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		*Alias
	}{ 
		GranularityType: o.GranularityType,
		
		Points: o.Points,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		Rank: o.Rank,
		Alias:    (*Alias)(o),
	})
}

func (o *Userbestpointsitem) UnmarshalJSON(b []byte) error {
	var UserbestpointsitemMap map[string]interface{}
	err := json.Unmarshal(b, &UserbestpointsitemMap)
	if err != nil {
		return err
	}
	
	if GranularityType, ok := UserbestpointsitemMap["granularityType"].(string); ok {
		o.GranularityType = &GranularityType
	}
    
	if Points, ok := UserbestpointsitemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if dateStartWorkdayString, ok := UserbestpointsitemMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := UserbestpointsitemMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if Rank, ok := UserbestpointsitemMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
