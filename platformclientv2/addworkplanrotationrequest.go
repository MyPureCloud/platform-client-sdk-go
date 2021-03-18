package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Addworkplanrotationrequest
type Addworkplanrotationrequest struct { 
	// Name - Name of this work plan rotation
	Name *string `json:"name,omitempty"`


	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Agents - Agents in this work plan rotation
	Agents *[]Addworkplanrotationagentrequest `json:"agents,omitempty"`


	// Pattern - Pattern with list of work plan IDs that rotate on a weekly basis
	Pattern *Workplanpatternrequest `json:"pattern,omitempty"`

}

// String returns a JSON representation of the model
func (o *Addworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
