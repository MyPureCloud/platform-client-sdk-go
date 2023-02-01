package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactioninput
type Updateactioninput struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`

	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`

	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`

	// Version - Version of this action
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateactioninput) SetField(field string, fieldValue interface{}) {
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

func (o Updateactioninput) MarshalJSON() ([]byte, error) {
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
	type Alias Updateactioninput
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		Category: o.Category,
		
		Name: o.Name,
		
		Config: o.Config,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Updateactioninput) UnmarshalJSON(b []byte) error {
	var UpdateactioninputMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateactioninputMap)
	if err != nil {
		return err
	}
	
	if Category, ok := UpdateactioninputMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Name, ok := UpdateactioninputMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Config, ok := UpdateactioninputMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if Version, ok := UpdateactioninputMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
