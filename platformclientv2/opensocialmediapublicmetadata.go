package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Opensocialmediapublicmetadata - Information about a public message.
type Opensocialmediapublicmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RootId - The id of the root public message.
	RootId *string `json:"rootId,omitempty"`

	// ReplyToId - The id of the message this public message is replying to.
	ReplyToId *string `json:"replyToId,omitempty"`

	// Source - The source of the public message. Useful when there could be more than location. Channel specific, e.g., for Facebook it's a source page.
	Source *string `json:"source,omitempty"`

	// Url - The URL of the social post on the native platform.
	Url *string `json:"url,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Opensocialmediapublicmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Opensocialmediapublicmetadata) MarshalJSON() ([]byte, error) {
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
	type Alias Opensocialmediapublicmetadata
	
	return json.Marshal(&struct { 
		RootId *string `json:"rootId,omitempty"`
		
		ReplyToId *string `json:"replyToId,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Url *string `json:"url,omitempty"`
		Alias
	}{ 
		RootId: o.RootId,
		
		ReplyToId: o.ReplyToId,
		
		Source: o.Source,
		
		Url: o.Url,
		Alias:    (Alias)(o),
	})
}

func (o *Opensocialmediapublicmetadata) UnmarshalJSON(b []byte) error {
	var OpensocialmediapublicmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &OpensocialmediapublicmetadataMap)
	if err != nil {
		return err
	}
	
	if RootId, ok := OpensocialmediapublicmetadataMap["rootId"].(string); ok {
		o.RootId = &RootId
	}
    
	if ReplyToId, ok := OpensocialmediapublicmetadataMap["replyToId"].(string); ok {
		o.ReplyToId = &ReplyToId
	}
    
	if Source, ok := OpensocialmediapublicmetadataMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Url, ok := OpensocialmediapublicmetadataMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Opensocialmediapublicmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
