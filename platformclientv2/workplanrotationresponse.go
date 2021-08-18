package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanrotationresponse
type Workplanrotationresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled - Whether the work plan rotation is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`


	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Pattern - Pattern with ordered list of work plans that rotate on a weekly basis
	Pattern *Workplanpatternresponse `json:"pattern,omitempty"`


	// AgentCount - Number of agents in this work plan rotation
	AgentCount *int `json:"agentCount,omitempty"`


	// Agents - Agents in this work plan rotation. Populate with expand=agents for GET WorkPlanRotationsList (defaults to empty list)
	Agents *[]Workplanrotationagentresponse `json:"agents,omitempty"`


	// Metadata - Version metadata for this work plan rotation
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Workplanrotationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanrotationresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Pattern *Workplanpatternresponse `json:"pattern,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		Agents *[]Workplanrotationagentresponse `json:"agents,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Enabled: u.Enabled,
		
		DateRange: u.DateRange,
		
		Pattern: u.Pattern,
		
		AgentCount: u.AgentCount,
		
		Agents: u.Agents,
		
		Metadata: u.Metadata,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workplanrotationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
