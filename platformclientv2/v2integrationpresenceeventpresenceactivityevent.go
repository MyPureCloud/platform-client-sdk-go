package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2integrationpresenceeventpresenceactivityevent
type V2integrationpresenceeventpresenceactivityevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId
	UserId *string `json:"userId,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// Source
	Source *string `json:"source,omitempty"`

	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// PresenceDefinition
	PresenceDefinition *V2integrationpresenceeventorganizationpresence `json:"presenceDefinition,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// CurrentDisplaySourceId
	CurrentDisplaySourceId *string `json:"currentDisplaySourceId,omitempty"`

	// PreviousDisplaySourceId
	PreviousDisplaySourceId *string `json:"previousDisplaySourceId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2integrationpresenceeventpresenceactivityevent) SetField(field string, fieldValue interface{}) {
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

func (o V2integrationpresenceeventpresenceactivityevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate", }
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
	type Alias V2integrationpresenceeventpresenceactivityevent
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		PresenceDefinition *V2integrationpresenceeventorganizationpresence `json:"presenceDefinition,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		CurrentDisplaySourceId *string `json:"currentDisplaySourceId,omitempty"`
		
		PreviousDisplaySourceId *string `json:"previousDisplaySourceId,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		EventType: o.EventType,
		
		Source: o.Source,
		
		ModifiedDate: ModifiedDate,
		
		PresenceDefinition: o.PresenceDefinition,
		
		Message: o.Message,
		
		CurrentDisplaySourceId: o.CurrentDisplaySourceId,
		
		PreviousDisplaySourceId: o.PreviousDisplaySourceId,
		Alias:    (Alias)(o),
	})
}

func (o *V2integrationpresenceeventpresenceactivityevent) UnmarshalJSON(b []byte) error {
	var V2integrationpresenceeventpresenceactivityeventMap map[string]interface{}
	err := json.Unmarshal(b, &V2integrationpresenceeventpresenceactivityeventMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := V2integrationpresenceeventpresenceactivityeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if EventType, ok := V2integrationpresenceeventpresenceactivityeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if Source, ok := V2integrationpresenceeventpresenceactivityeventMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if modifiedDateString, ok := V2integrationpresenceeventpresenceactivityeventMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if PresenceDefinition, ok := V2integrationpresenceeventpresenceactivityeventMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if Message, ok := V2integrationpresenceeventpresenceactivityeventMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if CurrentDisplaySourceId, ok := V2integrationpresenceeventpresenceactivityeventMap["currentDisplaySourceId"].(string); ok {
		o.CurrentDisplaySourceId = &CurrentDisplaySourceId
	}
    
	if PreviousDisplaySourceId, ok := V2integrationpresenceeventpresenceactivityeventMap["previousDisplaySourceId"].(string); ok {
		o.PreviousDisplaySourceId = &PreviousDisplaySourceId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2integrationpresenceeventpresenceactivityevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
