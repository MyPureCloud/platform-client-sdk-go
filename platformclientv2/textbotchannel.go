package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotchannel - Channel information relevant to a bot flow.
type Textbotchannel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the channel.
	Name *string `json:"name,omitempty"`

	// InputModes - The input modes for the channel.
	InputModes *[]string `json:"inputModes,omitempty"`

	// OutputModes - The output modes for the channel.
	OutputModes *[]string `json:"outputModes,omitempty"`

	// UserAgent - Information about the end user agent calling the bot flow.
	UserAgent *Textbotuseragent `json:"userAgent,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotchannel) SetField(field string, fieldValue interface{}) {
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

func (o Textbotchannel) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotchannel
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		InputModes *[]string `json:"inputModes,omitempty"`
		
		OutputModes *[]string `json:"outputModes,omitempty"`
		
		UserAgent *Textbotuseragent `json:"userAgent,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		InputModes: o.InputModes,
		
		OutputModes: o.OutputModes,
		
		UserAgent: o.UserAgent,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotchannel) UnmarshalJSON(b []byte) error {
	var TextbotchannelMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotchannelMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TextbotchannelMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InputModes, ok := TextbotchannelMap["inputModes"].([]interface{}); ok {
		InputModesString, _ := json.Marshal(InputModes)
		json.Unmarshal(InputModesString, &o.InputModes)
	}
	
	if OutputModes, ok := TextbotchannelMap["outputModes"].([]interface{}); ok {
		OutputModesString, _ := json.Marshal(OutputModes)
		json.Unmarshal(OutputModesString, &o.OutputModes)
	}
	
	if UserAgent, ok := TextbotchannelMap["userAgent"].(map[string]interface{}); ok {
		UserAgentString, _ := json.Marshal(UserAgent)
		json.Unmarshal(UserAgentString, &o.UserAgent)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
