package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Channel
type Channel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConnectUri
	ConnectUri *string `json:"connectUri,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// Expires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expires *time.Time `json:"expires,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Channel) SetField(field string, fieldValue interface{}) {
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

func (o Channel) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Expires", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Channel
	
	Expires := new(string)
	if o.Expires != nil {
		
		*Expires = timeutil.Strftime(o.Expires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expires = nil
	}
	
	return json.Marshal(&struct { 
		ConnectUri *string `json:"connectUri,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Expires *string `json:"expires,omitempty"`
		Alias
	}{ 
		ConnectUri: o.ConnectUri,
		
		Id: o.Id,
		
		Expires: Expires,
		Alias:    (Alias)(o),
	})
}

func (o *Channel) UnmarshalJSON(b []byte) error {
	var ChannelMap map[string]interface{}
	err := json.Unmarshal(b, &ChannelMap)
	if err != nil {
		return err
	}
	
	if ConnectUri, ok := ChannelMap["connectUri"].(string); ok {
		o.ConnectUri = &ConnectUri
	}
    
	if Id, ok := ChannelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if expiresString, ok := ChannelMap["expires"].(string); ok {
		Expires, _ := time.Parse("2006-01-02T15:04:05.999999Z", expiresString)
		o.Expires = &Expires
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Channel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
