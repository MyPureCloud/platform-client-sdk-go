package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationresponsesuggestionstopicsuggestedintent
type Conversationresponsesuggestionstopicsuggestedintent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Intent
	Intent *string `json:"intent,omitempty"`

	// IntentId
	IntentId *string `json:"intentId,omitempty"`

	// Confidence
	Confidence *float32 `json:"confidence,omitempty"`

	// DetectedSlots
	DetectedSlots *[]Conversationresponsesuggestionstopicsuggestedintentslot `json:"detectedSlots,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationresponsesuggestionstopicsuggestedintent) SetField(field string, fieldValue interface{}) {
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

func (o Conversationresponsesuggestionstopicsuggestedintent) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationresponsesuggestionstopicsuggestedintent
	
	return json.Marshal(&struct { 
		Intent *string `json:"intent,omitempty"`
		
		IntentId *string `json:"intentId,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		
		DetectedSlots *[]Conversationresponsesuggestionstopicsuggestedintentslot `json:"detectedSlots,omitempty"`
		Alias
	}{ 
		Intent: o.Intent,
		
		IntentId: o.IntentId,
		
		Confidence: o.Confidence,
		
		DetectedSlots: o.DetectedSlots,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationresponsesuggestionstopicsuggestedintent) UnmarshalJSON(b []byte) error {
	var ConversationresponsesuggestionstopicsuggestedintentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationresponsesuggestionstopicsuggestedintentMap)
	if err != nil {
		return err
	}
	
	if Intent, ok := ConversationresponsesuggestionstopicsuggestedintentMap["intent"].(string); ok {
		o.Intent = &Intent
	}
    
	if IntentId, ok := ConversationresponsesuggestionstopicsuggestedintentMap["intentId"].(string); ok {
		o.IntentId = &IntentId
	}
    
	if Confidence, ok := ConversationresponsesuggestionstopicsuggestedintentMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
    
	if DetectedSlots, ok := ConversationresponsesuggestionstopicsuggestedintentMap["detectedSlots"].([]interface{}); ok {
		DetectedSlotsString, _ := json.Marshal(DetectedSlots)
		json.Unmarshal(DetectedSlotsString, &o.DetectedSlots)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationresponsesuggestionstopicsuggestedintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
