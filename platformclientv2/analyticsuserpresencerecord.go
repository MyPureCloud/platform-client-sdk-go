package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserpresencerecord
type Analyticsuserpresencerecord struct { 
	// StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// SystemPresence - The user's system presence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// OrganizationPresenceId - The identifier for the user's organization presence
	OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`

}

func (o *Analyticsuserpresencerecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserpresencerecord
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`
		*Alias
	}{ 
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		SystemPresence: o.SystemPresence,
		
		OrganizationPresenceId: o.OrganizationPresenceId,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsuserpresencerecord) UnmarshalJSON(b []byte) error {
	var AnalyticsuserpresencerecordMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsuserpresencerecordMap)
	if err != nil {
		return err
	}
	
	if startTimeString, ok := AnalyticsuserpresencerecordMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := AnalyticsuserpresencerecordMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if SystemPresence, ok := AnalyticsuserpresencerecordMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if OrganizationPresenceId, ok := AnalyticsuserpresencerecordMap["organizationPresenceId"].(string); ok {
		o.OrganizationPresenceId = &OrganizationPresenceId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsuserpresencerecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
