package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingcontentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Recordingcontentactions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`

	// UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`

	// Textback - Text to be sent back in reply when the item is selected.
	Textback *string `json:"textback,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingcontentactions) SetField(field string, fieldValue interface{}) {
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

func (o Recordingcontentactions) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingcontentactions
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UrlTarget *string `json:"urlTarget,omitempty"`
		
		Textback *string `json:"textback,omitempty"`
		Alias
	}{ 
		Url: o.Url,
		
		UrlTarget: o.UrlTarget,
		
		Textback: o.Textback,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingcontentactions) UnmarshalJSON(b []byte) error {
	var RecordingcontentactionsMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingcontentactionsMap)
	if err != nil {
		return err
	}
	
	if Url, ok := RecordingcontentactionsMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if UrlTarget, ok := RecordingcontentactionsMap["urlTarget"].(string); ok {
		o.UrlTarget = &UrlTarget
	}
    
	if Textback, ok := RecordingcontentactionsMap["textback"].(string); ok {
		o.Textback = &Textback
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingcontentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
