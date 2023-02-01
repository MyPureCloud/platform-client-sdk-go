package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationblockedwebactionoffermessage
type Journeywebactioneventsnotificationblockedwebactionoffermessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action
	Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`

	// ActionMap
	ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`

	// ActionTarget
	ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`

	// BlockingReason
	BlockingReason *string `json:"blockingReason,omitempty"`

	// BlockingActionMap
	BlockingActionMap *Journeywebactioneventsnotificationactionmap `json:"blockingActionMap,omitempty"`

	// BlockingAction
	BlockingAction *Journeywebactioneventsnotificationeventaction `json:"blockingAction,omitempty"`

	// BlockingFrequencyCapBehaviour
	BlockingFrequencyCapBehaviour *string `json:"blockingFrequencyCapBehaviour,omitempty"`

	// BlockingPageUrlConditions
	BlockingPageUrlConditions *[]Journeywebactioneventsnotificationactionmappageurlcondition `json:"blockingPageUrlConditions,omitempty"`

	// BlockingScheduleGroup
	BlockingScheduleGroup *Journeywebactioneventsnotificationschedulegroup `json:"blockingScheduleGroup,omitempty"`

	// BlockingEmergencyScheduleGroup
	BlockingEmergencyScheduleGroup *Journeywebactioneventsnotificationemergencygroup `json:"blockingEmergencyScheduleGroup,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebactioneventsnotificationblockedwebactionoffermessage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeywebactioneventsnotificationblockedwebactionoffermessage
	
	return json.Marshal(&struct { 
		Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`
		
		ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		BlockingReason *string `json:"blockingReason,omitempty"`
		
		BlockingActionMap *Journeywebactioneventsnotificationactionmap `json:"blockingActionMap,omitempty"`
		
		BlockingAction *Journeywebactioneventsnotificationeventaction `json:"blockingAction,omitempty"`
		
		BlockingFrequencyCapBehaviour *string `json:"blockingFrequencyCapBehaviour,omitempty"`
		
		BlockingPageUrlConditions *[]Journeywebactioneventsnotificationactionmappageurlcondition `json:"blockingPageUrlConditions,omitempty"`
		
		BlockingScheduleGroup *Journeywebactioneventsnotificationschedulegroup `json:"blockingScheduleGroup,omitempty"`
		
		BlockingEmergencyScheduleGroup *Journeywebactioneventsnotificationemergencygroup `json:"blockingEmergencyScheduleGroup,omitempty"`
		Alias
	}{ 
		Action: o.Action,
		
		ActionMap: o.ActionMap,
		
		ActionTarget: o.ActionTarget,
		
		BlockingReason: o.BlockingReason,
		
		BlockingActionMap: o.BlockingActionMap,
		
		BlockingAction: o.BlockingAction,
		
		BlockingFrequencyCapBehaviour: o.BlockingFrequencyCapBehaviour,
		
		BlockingPageUrlConditions: o.BlockingPageUrlConditions,
		
		BlockingScheduleGroup: o.BlockingScheduleGroup,
		
		BlockingEmergencyScheduleGroup: o.BlockingEmergencyScheduleGroup,
		Alias:    (Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationblockedwebactionoffermessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationblockedwebactionoffermessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionMap, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ActionTarget, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if BlockingReason, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingReason"].(string); ok {
		o.BlockingReason = &BlockingReason
	}
    
	if BlockingActionMap, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingActionMap"].(map[string]interface{}); ok {
		BlockingActionMapString, _ := json.Marshal(BlockingActionMap)
		json.Unmarshal(BlockingActionMapString, &o.BlockingActionMap)
	}
	
	if BlockingAction, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingAction"].(map[string]interface{}); ok {
		BlockingActionString, _ := json.Marshal(BlockingAction)
		json.Unmarshal(BlockingActionString, &o.BlockingAction)
	}
	
	if BlockingFrequencyCapBehaviour, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingFrequencyCapBehaviour"].(string); ok {
		o.BlockingFrequencyCapBehaviour = &BlockingFrequencyCapBehaviour
	}
    
	if BlockingPageUrlConditions, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingPageUrlConditions"].([]interface{}); ok {
		BlockingPageUrlConditionsString, _ := json.Marshal(BlockingPageUrlConditions)
		json.Unmarshal(BlockingPageUrlConditionsString, &o.BlockingPageUrlConditions)
	}
	
	if BlockingScheduleGroup, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingScheduleGroup"].(map[string]interface{}); ok {
		BlockingScheduleGroupString, _ := json.Marshal(BlockingScheduleGroup)
		json.Unmarshal(BlockingScheduleGroupString, &o.BlockingScheduleGroup)
	}
	
	if BlockingEmergencyScheduleGroup, ok := JourneywebactioneventsnotificationblockedwebactionoffermessageMap["blockingEmergencyScheduleGroup"].(map[string]interface{}); ok {
		BlockingEmergencyScheduleGroupString, _ := json.Marshal(BlockingEmergencyScheduleGroup)
		json.Unmarshal(BlockingEmergencyScheduleGroupString, &o.BlockingEmergencyScheduleGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationblockedwebactionoffermessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
