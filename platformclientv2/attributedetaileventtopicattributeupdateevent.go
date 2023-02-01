package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Attributedetaileventtopicattributeupdateevent
type Attributedetaileventtopicattributeupdateevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`

	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Attributedetaileventtopicattributeupdateevent) SetField(field string, fieldValue interface{}) {
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

func (o Attributedetaileventtopicattributeupdateevent) MarshalJSON() ([]byte, error) {
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
	type Alias Attributedetaileventtopicattributeupdateevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		Attributes: o.Attributes,
		Alias:    (Alias)(o),
	})
}

func (o *Attributedetaileventtopicattributeupdateevent) UnmarshalJSON(b []byte) error {
	var AttributedetaileventtopicattributeupdateeventMap map[string]interface{}
	err := json.Unmarshal(b, &AttributedetaileventtopicattributeupdateeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := AttributedetaileventtopicattributeupdateeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := AttributedetaileventtopicattributeupdateeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := AttributedetaileventtopicattributeupdateeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if Attributes, ok := AttributedetaileventtopicattributeupdateeventMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attributedetaileventtopicattributeupdateevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
