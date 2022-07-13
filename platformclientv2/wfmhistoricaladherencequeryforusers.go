package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencequeryforusers
type Wfmhistoricaladherencequeryforusers struct { 
	// StartDate - Beginning of the date range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
	EndDate *time.Time `json:"endDate,omitempty"`


	// TimeZone - The time zone, in olson format, to use in defining days when computing adherence. If it is not set, the business unit time zone will be used. The results will be returned as UTC timestamps regardless of the time zone input.
	TimeZone *string `json:"timeZone,omitempty"`


	// UserIds - The userIds to report on
	UserIds *[]string `json:"userIds,omitempty"`


	// IncludeExceptions - Whether user exceptions should be returned as part of the results
	IncludeExceptions *bool `json:"includeExceptions,omitempty"`

}

func (o *Wfmhistoricaladherencequeryforusers) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencequeryforusers
	
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
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		TimeZone: o.TimeZone,
		
		UserIds: o.UserIds,
		
		IncludeExceptions: o.IncludeExceptions,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencequeryforusers) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencequeryforusersMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencequeryforusersMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmhistoricaladherencequeryforusersMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmhistoricaladherencequeryforusersMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if TimeZone, ok := WfmhistoricaladherencequeryforusersMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if UserIds, ok := WfmhistoricaladherencequeryforusersMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if IncludeExceptions, ok := WfmhistoricaladherencequeryforusersMap["includeExceptions"].(bool); ok {
		o.IncludeExceptions = &IncludeExceptions
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencequeryforusers) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
