package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactivitycoderequest
type Updateactivitycoderequest struct { 
	// Name - The name of the activity code
	Name *string `json:"name,omitempty"`


	// Category - The activity code's category. Attempting to change the category of a default activity code will return an error
	Category *string `json:"category,omitempty"`


	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsWorkTime - Indicates whether or not the activity should be counted as work time
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`


	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`


	// CountsTowardShrinkage - Whether or not this activity code counts toward shrinkage calculations
	CountsTowardShrinkage *bool `json:"countsTowardShrinkage,omitempty"`


	// PlannedShrinkage - Whether this activity code is considered planned or unplanned shrinkage
	PlannedShrinkage *bool `json:"plannedShrinkage,omitempty"`


	// Interruptible - Whether this activity code is considered interruptible
	Interruptible *bool `json:"interruptible,omitempty"`


	// SecondaryPresences - The secondary presences of this activity code
	SecondaryPresences *Listwrappersecondarypresence `json:"secondaryPresences,omitempty"`


	// Metadata - Version metadata for the associated business unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updateactivitycoderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateactivitycoderequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`
		
		AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`
		
		CountsTowardShrinkage *bool `json:"countsTowardShrinkage,omitempty"`
		
		PlannedShrinkage *bool `json:"plannedShrinkage,omitempty"`
		
		Interruptible *bool `json:"interruptible,omitempty"`
		
		SecondaryPresences *Listwrappersecondarypresence `json:"secondaryPresences,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
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
		Alias:    (*Alias)(o),
	})
}

func (o *Updateactivitycoderequest) UnmarshalJSON(b []byte) error {
	var UpdateactivitycoderequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateactivitycoderequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateactivitycoderequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Category, ok := UpdateactivitycoderequestMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if LengthInMinutes, ok := UpdateactivitycoderequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if CountsAsPaidTime, ok := UpdateactivitycoderequestMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if CountsAsWorkTime, ok := UpdateactivitycoderequestMap["countsAsWorkTime"].(bool); ok {
		o.CountsAsWorkTime = &CountsAsWorkTime
	}
    
	if AgentTimeOffSelectable, ok := UpdateactivitycoderequestMap["agentTimeOffSelectable"].(bool); ok {
		o.AgentTimeOffSelectable = &AgentTimeOffSelectable
	}
    
	if CountsTowardShrinkage, ok := UpdateactivitycoderequestMap["countsTowardShrinkage"].(bool); ok {
		o.CountsTowardShrinkage = &CountsTowardShrinkage
	}
    
	if PlannedShrinkage, ok := UpdateactivitycoderequestMap["plannedShrinkage"].(bool); ok {
		o.PlannedShrinkage = &PlannedShrinkage
	}
    
	if Interruptible, ok := UpdateactivitycoderequestMap["interruptible"].(bool); ok {
		o.Interruptible = &Interruptible
	}
    
	if SecondaryPresences, ok := UpdateactivitycoderequestMap["secondaryPresences"].(map[string]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	
	if Metadata, ok := UpdateactivitycoderequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateactivitycoderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
