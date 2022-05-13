package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Currentuserschedulerequestbody - POST request body for fetching the current user's schedule over a given range
type Currentuserschedulerequestbody struct { 
	// StartDate - Beginning of the range of schedules to fetch, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the range of schedules to fetch, in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// LoadFullWeeks - Whether to load the full week's schedule (for the current user) of any week overlapping the start/end date query parameters, defaults to false
	LoadFullWeeks *bool `json:"loadFullWeeks,omitempty"`

}

func (o *Currentuserschedulerequestbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Currentuserschedulerequestbody
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		LoadFullWeeks *bool `json:"loadFullWeeks,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		LoadFullWeeks: o.LoadFullWeeks,
		Alias:    (*Alias)(o),
	})
}

func (o *Currentuserschedulerequestbody) UnmarshalJSON(b []byte) error {
	var CurrentuserschedulerequestbodyMap map[string]interface{}
	err := json.Unmarshal(b, &CurrentuserschedulerequestbodyMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := CurrentuserschedulerequestbodyMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := CurrentuserschedulerequestbodyMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if LoadFullWeeks, ok := CurrentuserschedulerequestbodyMap["loadFullWeeks"].(bool); ok {
		o.LoadFullWeeks = &LoadFullWeeks
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Currentuserschedulerequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
