package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentconfigurationversionentityref
type Webdeploymentconfigurationversionentityref struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The configuration version ID
	Id *string `json:"id,omitempty"`

	// Name - The configuration version name
	Name *string `json:"name,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// Version - The version of the configuration
	Version *string `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webdeploymentconfigurationversionentityref) SetField(field string, fieldValue interface{}) {
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

func (o Webdeploymentconfigurationversionentityref) MarshalJSON() ([]byte, error) {
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
	type Alias Webdeploymentconfigurationversionentityref
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Version *string `json:"version,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Webdeploymentconfigurationversionentityref) UnmarshalJSON(b []byte) error {
	var WebdeploymentconfigurationversionentityrefMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentconfigurationversionentityrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentconfigurationversionentityrefMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebdeploymentconfigurationversionentityrefMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := WebdeploymentconfigurationversionentityrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Version, ok := WebdeploymentconfigurationversionentityrefMap["version"].(string); ok {
		o.Version = &Version
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
