package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateworkplanrotationrequest
type Updateworkplanrotationrequest struct { 
	// Name - Name of this work plan rotation
	Name *string `json:"name,omitempty"`


	// Enabled - Whether the work plan rotation is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`


	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Agents - Agents in this work plan rotation
	Agents *[]Updateworkplanrotationagentrequest `json:"agents,omitempty"`


	// Pattern - Pattern with list of work plan IDs that rotate on a weekly basis
	Pattern *Workplanpatternrequest `json:"pattern,omitempty"`


	// Metadata - Version metadata for this work plan rotation
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Updateworkplanrotationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateworkplanrotationrequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Agents *[]Updateworkplanrotationagentrequest `json:"agents,omitempty"`
		
		Pattern *Workplanpatternrequest `json:"pattern,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Enabled: u.Enabled,
		
		DateRange: u.DateRange,
		
		Agents: u.Agents,
		
		Pattern: u.Pattern,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
