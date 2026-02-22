package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Customattributespatchrequest
type Customattributespatchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique identifier for the Custom Attributes record. IDs are created by users.
	Id *string `json:"id,omitempty"`

	// Divisions - The list of division ids. Use [] if divisions aren't used (Unassigned Division). Omitting or setting to [] clears existing values on update.
	Divisions *[]string `json:"divisions,omitempty"`

	// Version - The latest version of the Custom Attributes record. Optional for concurrency check.
	Version *int `json:"version,omitempty"`

	// CustomAttributes - The map of attribute values.
	CustomAttributes *map[string]interface{} `json:"customAttributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Customattributespatchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Customattributespatchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Customattributespatchrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Divisions *[]string `json:"divisions,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		CustomAttributes *map[string]interface{} `json:"customAttributes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Divisions: o.Divisions,
		
		Version: o.Version,
		
		CustomAttributes: o.CustomAttributes,
		Alias:    (Alias)(o),
	})
}

func (o *Customattributespatchrequest) UnmarshalJSON(b []byte) error {
	var CustomattributespatchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CustomattributespatchrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CustomattributespatchrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Divisions, ok := CustomattributespatchrequestMap["divisions"].([]interface{}); ok {
		DivisionsString, _ := json.Marshal(Divisions)
		json.Unmarshal(DivisionsString, &o.Divisions)
	}
	
	if Version, ok := CustomattributespatchrequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CustomAttributes, ok := CustomattributespatchrequestMap["customAttributes"].(map[string]interface{}); ok {
		CustomAttributesString, _ := json.Marshal(CustomAttributes)
		json.Unmarshal(CustomAttributesString, &o.CustomAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Customattributespatchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
