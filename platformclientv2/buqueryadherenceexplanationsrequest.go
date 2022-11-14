package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buqueryadherenceexplanationsrequest
type Buqueryadherenceexplanationsrequest struct { 
	// StartDate - The start date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the range to query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// ManagementUnitIds - A filter for which management units to query. Leave empty or omit entirely for all management units in the business unit
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`


	// AgentIds - A filter for which agents within the business unit to query. Leave empty or omit entirely for all agents in the business unit (or management units if specified)
	AgentIds *[]string `json:"agentIds,omitempty"`

}

func (o *Buqueryadherenceexplanationsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buqueryadherenceexplanationsrequest
	
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
		
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		
		AgentIds *[]string `json:"agentIds,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		ManagementUnitIds: o.ManagementUnitIds,
		
		AgentIds: o.AgentIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buqueryadherenceexplanationsrequest) UnmarshalJSON(b []byte) error {
	var BuqueryadherenceexplanationsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BuqueryadherenceexplanationsrequestMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuqueryadherenceexplanationsrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BuqueryadherenceexplanationsrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if ManagementUnitIds, ok := BuqueryadherenceexplanationsrequestMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	
	if AgentIds, ok := BuqueryadherenceexplanationsrequestMap["agentIds"].([]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buqueryadherenceexplanationsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
