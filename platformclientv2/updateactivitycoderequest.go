package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactivitycoderequest
type Updateactivitycoderequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// PlanningGroupIds - The planning group IDs associated with this activity code
	PlanningGroupIds *Listwrapperstring `json:"planningGroupIds,omitempty"`

	// Metadata - Version metadata for the associated business unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateactivitycoderequest) SetField(field string, fieldValue interface{}) {
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

func (o Updateactivitycoderequest) MarshalJSON() ([]byte, error) {
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
		
		PlanningGroupIds *Listwrapperstring `json:"planningGroupIds,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
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
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
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
	
	if PlanningGroupIds, ok := UpdateactivitycoderequestMap["planningGroupIds"].(map[string]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
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
