package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentlocation - Location object.
type Conversationcontentlocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url - URL of the Location.
	Url *string `json:"url,omitempty"`

	// Address - Location postal address.
	Address *string `json:"address,omitempty"`

	// Text - Location name.
	Text *string `json:"text,omitempty"`

	// Latitude - Latitude of the location.
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude - Longitude of the location.
	Longitude *float64 `json:"longitude,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentlocation) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentlocation) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentlocation
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Latitude *float64 `json:"latitude,omitempty"`
		
		Longitude *float64 `json:"longitude,omitempty"`
		Alias
	}{ 
		Url: o.Url,
		
		Address: o.Address,
		
		Text: o.Text,
		
		Latitude: o.Latitude,
		
		Longitude: o.Longitude,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentlocation) UnmarshalJSON(b []byte) error {
	var ConversationcontentlocationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentlocationMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationcontentlocationMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Address, ok := ConversationcontentlocationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Text, ok := ConversationcontentlocationMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Latitude, ok := ConversationcontentlocationMap["latitude"].(float64); ok {
		o.Latitude = &Latitude
	}
    
	if Longitude, ok := ConversationcontentlocationMap["longitude"].(float64); ok {
		o.Longitude = &Longitude
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}