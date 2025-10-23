package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createactivitycoderequest
type Createactivitycoderequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// PlanningGroupIds - The planning group IDs associated with this activity code
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createactivitycoderequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Createactivitycoderequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		Alias
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
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
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
	
	if PlanningGroupIds, ok := CreateactivitycoderequestMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createactivitycoderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
