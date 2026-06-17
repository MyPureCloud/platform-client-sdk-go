package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentschedulebiddingpreferenceresponse
type Agentschedulebiddingpreferenceresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Submitted - Whether the preference is submitted
	Submitted *bool `json:"submitted,omitempty"`

	// AssignedScheduleSetId - The schedule set assigned to the agent by the bid process. Will be set after bid is processed
	AssignedScheduleSetId *string `json:"assignedScheduleSetId,omitempty"`

	// OverriddenScheduleSetId - The schedule set that overrides the assigned schedule set for the agent
	OverriddenScheduleSetId *string `json:"overriddenScheduleSetId,omitempty"`

	// OverrideReason - The reason why the assigned schedule set has been overridden. This must be null without an override schedule set
	OverrideReason *string `json:"overrideReason,omitempty"`

	// AgentScheduleBidPreferences - The schedule bidding preferences
	AgentScheduleBidPreferences *[]Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentschedulebiddingpreferenceresponse) SetField(field string, fieldValue interface{}) {
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

func (o Agentschedulebiddingpreferenceresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Agentschedulebiddingpreferenceresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Submitted *bool `json:"submitted,omitempty"`
		
		AssignedScheduleSetId *string `json:"assignedScheduleSetId,omitempty"`
		
		OverriddenScheduleSetId *string `json:"overriddenScheduleSetId,omitempty"`
		
		OverrideReason *string `json:"overrideReason,omitempty"`
		
		AgentScheduleBidPreferences *[]Agentschedulebiddingpreferencepriority `json:"agentScheduleBidPreferences,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Submitted: o.Submitted,
		
		AssignedScheduleSetId: o.AssignedScheduleSetId,
		
		OverriddenScheduleSetId: o.OverriddenScheduleSetId,
		
		OverrideReason: o.OverrideReason,
		
		AgentScheduleBidPreferences: o.AgentScheduleBidPreferences,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentschedulebiddingpreferenceresponse) UnmarshalJSON(b []byte) error {
	var AgentschedulebiddingpreferenceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AgentschedulebiddingpreferenceresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentschedulebiddingpreferenceresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Submitted, ok := AgentschedulebiddingpreferenceresponseMap["submitted"].(bool); ok {
		o.Submitted = &Submitted
	}
    
	if AssignedScheduleSetId, ok := AgentschedulebiddingpreferenceresponseMap["assignedScheduleSetId"].(string); ok {
		o.AssignedScheduleSetId = &AssignedScheduleSetId
	}
    
	if OverriddenScheduleSetId, ok := AgentschedulebiddingpreferenceresponseMap["overriddenScheduleSetId"].(string); ok {
		o.OverriddenScheduleSetId = &OverriddenScheduleSetId
	}
    
	if OverrideReason, ok := AgentschedulebiddingpreferenceresponseMap["overrideReason"].(string); ok {
		o.OverrideReason = &OverrideReason
	}
    
	if AgentScheduleBidPreferences, ok := AgentschedulebiddingpreferenceresponseMap["agentScheduleBidPreferences"].([]interface{}); ok {
		AgentScheduleBidPreferencesString, _ := json.Marshal(AgentScheduleBidPreferences)
		json.Unmarshal(AgentScheduleBidPreferencesString, &o.AgentScheduleBidPreferences)
	}
	
	if SelfUri, ok := AgentschedulebiddingpreferenceresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentschedulebiddingpreferenceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
