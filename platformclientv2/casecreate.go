package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Casecreate
type Casecreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CaseplanId - The ID of the caseplan to create the case from.
	CaseplanId *string `json:"caseplanId,omitempty"`

	// OwnerId - The ID of the owner of the case.
	OwnerId *string `json:"ownerId,omitempty"`

	// Summary - Overview information for the Case. Valid length between 3 and 512 characters.
	Summary *string `json:"summary,omitempty"`

	// ExternalContactId - The ID of the External Contact associated with the Case.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ConversationId - The ID of conversation associated with the Case.
	ConversationId *string `json:"conversationId,omitempty"`

	// WorkitemId - The ID of the workitem associated with the Case.
	WorkitemId *string `json:"workitemId,omitempty"`

	// TtlSeconds - The epoch timestamp in seconds specifying the time-to-live for the lifetime of the Case. Can not be greater than 365 days from the current time.
	TtlSeconds *int `json:"ttlSeconds,omitempty"`

	// Intake - The intake data for the Case. Maximum of 10 intake objects allowed.
	Intake *[]Intake `json:"intake,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Casecreate) SetField(field string, fieldValue interface{}) {
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

func (o Casecreate) MarshalJSON() ([]byte, error) {
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
	type Alias Casecreate
	
	return json.Marshal(&struct { 
		CaseplanId *string `json:"caseplanId,omitempty"`
		
		OwnerId *string `json:"ownerId,omitempty"`
		
		Summary *string `json:"summary,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		WorkitemId *string `json:"workitemId,omitempty"`
		
		TtlSeconds *int `json:"ttlSeconds,omitempty"`
		
		Intake *[]Intake `json:"intake,omitempty"`
		Alias
	}{ 
		CaseplanId: o.CaseplanId,
		
		OwnerId: o.OwnerId,
		
		Summary: o.Summary,
		
		ExternalContactId: o.ExternalContactId,
		
		ConversationId: o.ConversationId,
		
		WorkitemId: o.WorkitemId,
		
		TtlSeconds: o.TtlSeconds,
		
		Intake: o.Intake,
		Alias:    (Alias)(o),
	})
}

func (o *Casecreate) UnmarshalJSON(b []byte) error {
	var CasecreateMap map[string]interface{}
	err := json.Unmarshal(b, &CasecreateMap)
	if err != nil {
		return err
	}
	
	if CaseplanId, ok := CasecreateMap["caseplanId"].(string); ok {
		o.CaseplanId = &CaseplanId
	}
    
	if OwnerId, ok := CasecreateMap["ownerId"].(string); ok {
		o.OwnerId = &OwnerId
	}
    
	if Summary, ok := CasecreateMap["summary"].(string); ok {
		o.Summary = &Summary
	}
    
	if ExternalContactId, ok := CasecreateMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ConversationId, ok := CasecreateMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if WorkitemId, ok := CasecreateMap["workitemId"].(string); ok {
		o.WorkitemId = &WorkitemId
	}
    
	if TtlSeconds, ok := CasecreateMap["ttlSeconds"].(float64); ok {
		TtlSecondsInt := int(TtlSeconds)
		o.TtlSeconds = &TtlSecondsInt
	}
	
	if Intake, ok := CasecreateMap["intake"].([]interface{}); ok {
		IntakeString, _ := json.Marshal(Intake)
		json.Unmarshal(IntakeString, &o.Intake)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Casecreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
