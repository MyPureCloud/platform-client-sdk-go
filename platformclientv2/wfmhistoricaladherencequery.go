package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencequery
type Wfmhistoricaladherencequery struct { 
	// StartDate - Beginning of the date range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
	EndDate *time.Time `json:"endDate,omitempty"`


	// TimeZone - The time zone to use for returned results in olson format. If it is not set, the business unit time zone will be used to compute adherence
	TimeZone *string `json:"timeZone,omitempty"`


	// UserIds - The userIds to report on. If null or not set, adherence will be computed for all the users in management unit or requested teamIds
	UserIds *[]string `json:"userIds,omitempty"`


	// IncludeExceptions - Whether user exceptions should be returned as part of the results
	IncludeExceptions *bool `json:"includeExceptions,omitempty"`


	// TeamIds - The teamIds to report on. If null or not set, adherence will be computed for requested users if applicable or otherwise all users in the management unit. Note: If teamIds is also specified, only adherence for users in the requested teams will be returned
	TeamIds *[]string `json:"teamIds,omitempty"`

}

func (o *Wfmhistoricaladherencequery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencequery
	
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
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		IncludeExceptions *bool `json:"includeExceptions,omitempty"`
		
		TeamIds *[]string `json:"teamIds,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		TimeZone: o.TimeZone,
		
		UserIds: o.UserIds,
		
		IncludeExceptions: o.IncludeExceptions,
		
		TeamIds: o.TeamIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencequery) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencequeryMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencequeryMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmhistoricaladherencequeryMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmhistoricaladherencequeryMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if TimeZone, ok := WfmhistoricaladherencequeryMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if UserIds, ok := WfmhistoricaladherencequeryMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if IncludeExceptions, ok := WfmhistoricaladherencequeryMap["includeExceptions"].(bool); ok {
		o.IncludeExceptions = &IncludeExceptions
	}
    
	if TeamIds, ok := WfmhistoricaladherencequeryMap["teamIds"].([]interface{}); ok {
		TeamIdsString, _ := json.Marshal(TeamIds)
		json.Unmarshal(TeamIdsString, &o.TeamIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
