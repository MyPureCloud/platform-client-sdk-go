package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Updateworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
