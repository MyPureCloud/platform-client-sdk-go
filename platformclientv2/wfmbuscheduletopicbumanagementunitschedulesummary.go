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

func (u *Wfmbuscheduletopicbumanagementunitschedulesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbumanagementunitschedulesummary

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		ManagementUnit: u.ManagementUnit,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Agents: u.Agents,
		
		AgentCount: u.AgentCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
