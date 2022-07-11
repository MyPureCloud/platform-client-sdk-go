package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createactivitycoderequest
type Createactivitycoderequest struct { 
	// Name - The name of the activity code
	Name *string `json:"name,omitempty"`


	// Category - The activity code's category
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
	SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`

}

func (o *Createactivitycoderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createactivitycoderequest
	
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
		
		SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`
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
		Alias:    (*Alias)(o),
	})
}

func (o *Createactivitycoderequest) UnmarshalJSON(b []byte) error {
	var CreateactivitycoderequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateactivitycoderequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateactivitycoderequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Category, ok := CreateactivitycoderequestMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if LengthInMinutes, ok := CreateactivitycoderequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if CountsAsPaidTime, ok := CreateactivitycoderequestMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if CountsAsWorkTime, ok := CreateactivitycoderequestMap["countsAsWorkTime"].(bool); ok {
		o.CountsAsWorkTime = &CountsAsWorkTime
	}
    
	if AgentTimeOffSelectable, ok := CreateactivitycoderequestMap["agentTimeOffSelectable"].(bool); ok {
		o.AgentTimeOffSelectable = &AgentTimeOffSelectable
	}
    
	if CountsTowardShrinkage, ok := CreateactivitycoderequestMap["countsTowardShrinkage"].(bool); ok {
		o.CountsTowardShrinkage = &CountsTowardShrinkage
	}
    
	if PlannedShrinkage, ok := CreateactivitycoderequestMap["plannedShrinkage"].(bool); ok {
		o.PlannedShrinkage = &PlannedShrinkage
	}
    
	if Interruptible, ok := CreateactivitycoderequestMap["interruptible"].(bool); ok {
		o.Interruptible = &Interruptible
	}
    
	if SecondaryPresences, ok := CreateactivitycoderequestMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createactivitycoderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
