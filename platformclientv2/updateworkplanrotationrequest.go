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

func (o *Updateworkplanrotationrequest) MarshalJSON() ([]byte, error) {
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
		Name: o.Name,
		
		Enabled: o.Enabled,
		
		DateRange: o.DateRange,
		
		Agents: o.Agents,
		
		Pattern: o.Pattern,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateworkplanrotationrequest) UnmarshalJSON(b []byte) error {
	var UpdateworkplanrotationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateworkplanrotationrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateworkplanrotationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := UpdateworkplanrotationrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if DateRange, ok := UpdateworkplanrotationrequestMap["dateRange"].(map[string]interface{}); ok {
		DateRangeString, _ := json.Marshal(DateRange)
		json.Unmarshal(DateRangeString, &o.DateRange)
	}
	
	if Agents, ok := UpdateworkplanrotationrequestMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if Pattern, ok := UpdateworkplanrotationrequestMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if Metadata, ok := UpdateworkplanrotationrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
