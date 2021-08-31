package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbumanagementunitschedulesummary
type Wfmbuscheduletopicbumanagementunitschedulesummary struct { 
	// ManagementUnit
	ManagementUnit *Wfmbuscheduletopicmanagementunit `json:"managementUnit,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// Agents
	Agents *[]Wfmbuscheduletopicuserreference `json:"agents,omitempty"`


	// AgentCount
	AgentCount *int `json:"agentCount,omitempty"`

}

func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbumanagementunitschedulesummary
	
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
		ManagementUnit *Wfmbuscheduletopicmanagementunit `json:"managementUnit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Agents *[]Wfmbuscheduletopicuserreference `json:"agents,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		*Alias
	}{ 
		ManagementUnit: o.ManagementUnit,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Agents: o.Agents,
		
		AgentCount: o.AgentCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbumanagementunitschedulesummaryMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbumanagementunitschedulesummaryMap)
	if err != nil {
		return err
	}
	
	if ManagementUnit, ok := WfmbuscheduletopicbumanagementunitschedulesummaryMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if startDateString, ok := WfmbuscheduletopicbumanagementunitschedulesummaryMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmbuscheduletopicbumanagementunitschedulesummaryMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Agents, ok := WfmbuscheduletopicbumanagementunitschedulesummaryMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if AgentCount, ok := WfmbuscheduletopicbumanagementunitschedulesummaryMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
