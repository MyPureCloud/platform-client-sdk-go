package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterid - User information for a twitter account. Either id OR screenName (or both) must be present
type Twitterid struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - twitter user.id_str. Max: 255 characters
	Id *string `json:"id,omitempty"`

	// Name - twitter user.name. Max: 255 characters
	Name *string `json:"name,omitempty"`

	// ScreenName - twitter user.screen_name. Max: 255 characters. Must match pattern: ^@?[A-Za-z0-9_]+$
	ScreenName *string `json:"screenName,omitempty"`

	// Verified - whether this data has been verified using the twitter API
	Verified *bool `json:"verified,omitempty"`

	// ProfileUrl - url of user's twitter profile
	ProfileUrl *string `json:"profileUrl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Twitterid) SetField(field string, fieldValue interface{}) {
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

func (o Twitterid) MarshalJSON() ([]byte, error) {
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
	type Alias Twitterid
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ScreenName *string `json:"screenName,omitempty"`
		
		Verified *bool `json:"verified,omitempty"`
		
		ProfileUrl *string `json:"profileUrl,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ScreenName: o.ScreenName,
		
		Verified: o.Verified,
		
		ProfileUrl: o.ProfileUrl,
		Alias:    (Alias)(o),
	})
}

func (o *Twitterid) UnmarshalJSON(b []byte) error {
	var TwitteridMap map[string]interface{}
	err := json.Unmarshal(b, &TwitteridMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitteridMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TwitteridMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ScreenName, ok := TwitteridMap["screenName"].(string); ok {
		o.ScreenName = &ScreenName
	}
    
	if Verified, ok := TwitteridMap["verified"].(bool); ok {
		o.Verified = &Verified
	}
    
	if ProfileUrl, ok := TwitteridMap["profileUrl"].(string); ok {
		o.ProfileUrl = &ProfileUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
