package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitactivitycode
type Businessunitactivitycode struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Active - Whether this activity code is active or has been deleted
	Active *bool `json:"active,omitempty"`


	// DefaultCode - Whether this is a default activity code
	DefaultCode *bool `json:"defaultCode,omitempty"`


	// Category - The category of the activity code
	Category *string `json:"category,omitempty"`


	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsWorkTime - Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`


	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`


	// CountsTowardShrinkage - Whether or not this activity code counts toward shrinkage calculations
	CountsTowardShrinkage *bool `json:"countsTowardShrinkage,omitempty"`


	// PlannedShrinkage - Whether this activity code is considered planned or unplanned shrinkage
	PlannedShrinkage *bool `json:"plannedShrinkage,omitempty"`


	// Interruptible - Whether this activity code is considered interruptible
	Interruptible *bool `json:"interruptible,omitempty"`


	// SecondaryPresences - The secondary presences of this activity code
	SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`


	// Metadata - Version metadata of this activity code
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Businessunitactivitycode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Businessunitactivitycode
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		DefaultCode *bool `json:"defaultCode,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`
		
		AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`
		
		CountsTowardShrinkage *bool `json:"countsTowardShrinkage,omitempty"`
		
		PlannedShrinkage *bool `json:"plannedShrinkage,omitempty"`
		
		Interruptible *bool `json:"interruptible,omitempty"`
		
		SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Active: o.Active,
		
		DefaultCode: o.DefaultCode,
		
		Category: o.Category,
		
		LengthInMinutes: o.LengthInMinutes,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		CountsAsWorkTime: o.CountsAsWorkTime,
		
		AgentTimeOffSelectable: o.AgentTimeOffSelectable,
		
		CountsTowardShrinkage: o.CountsTowardShrinkage,
		
		PlannedShrinkage: o.PlannedShrinkage,
		
		Interruptible: o.Interruptible,
		
		SecondaryPresences: o.SecondaryPresences,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Businessunitactivitycode) UnmarshalJSON(b []byte) error {
	var BusinessunitactivitycodeMap map[string]interface{}
	err := json.Unmarshal(b, &BusinessunitactivitycodeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BusinessunitactivitycodeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BusinessunitactivitycodeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Active, ok := BusinessunitactivitycodeMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if DefaultCode, ok := BusinessunitactivitycodeMap["defaultCode"].(bool); ok {
		o.DefaultCode = &DefaultCode
	}
    
	if Category, ok := BusinessunitactivitycodeMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if LengthInMinutes, ok := BusinessunitactivitycodeMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if CountsAsPaidTime, ok := BusinessunitactivitycodeMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if CountsAsWorkTime, ok := BusinessunitactivitycodeMap["countsAsWorkTime"].(bool); ok {
		o.CountsAsWorkTime = &CountsAsWorkTime
	}
    
	if AgentTimeOffSelectable, ok := BusinessunitactivitycodeMap["agentTimeOffSelectable"].(bool); ok {
		o.AgentTimeOffSelectable = &AgentTimeOffSelectable
	}
    
	if CountsTowardShrinkage, ok := BusinessunitactivitycodeMap["countsTowardShrinkage"].(bool); ok {
		o.CountsTowardShrinkage = &CountsTowardShrinkage
	}
    
	if PlannedShrinkage, ok := BusinessunitactivitycodeMap["plannedShrinkage"].(bool); ok {
		o.PlannedShrinkage = &PlannedShrinkage
	}
    
	if Interruptible, ok := BusinessunitactivitycodeMap["interruptible"].(bool); ok {
		o.Interruptible = &Interruptible
	}
    
	if SecondaryPresences, ok := BusinessunitactivitycodeMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	
	if Metadata, ok := BusinessunitactivitycodeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := BusinessunitactivitycodeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Businessunitactivitycode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
