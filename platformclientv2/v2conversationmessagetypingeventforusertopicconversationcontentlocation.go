package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentlocation
type V2conversationmessagetypingeventforusertopicconversationcontentlocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url
	Url *string `json:"url,omitempty"`

	// Address
	Address *string `json:"address,omitempty"`

	// Text
	Text *string `json:"text,omitempty"`

	// Latitude
	Latitude *float32 `json:"latitude,omitempty"`

	// Longitude
	Longitude *float32 `json:"longitude,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforusertopicconversationcontentlocation) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentlocation
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Latitude *float32 `json:"latitude,omitempty"`
		
		Longitude *float32 `json:"longitude,omitempty"`
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

func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentlocationMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentlocationMap)
	if err != nil {
		return err
	}
	
	if Url, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Address, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Latitude, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["latitude"].(float64); ok {
		LatitudeFloat32 := float32(Latitude)
		o.Latitude = &LatitudeFloat32
	}
    
	if Longitude, ok := V2conversationmessagetypingeventforusertopicconversationcontentlocationMap["longitude"].(float64); ok {
		LongitudeFloat32 := float32(Longitude)
		o.Longitude = &LongitudeFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
