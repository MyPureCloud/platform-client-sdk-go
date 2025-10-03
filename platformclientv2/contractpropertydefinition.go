package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contractpropertydefinition
type Contractpropertydefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title
	Title *string `json:"title,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// VarType
	VarType *[]string `json:"type,omitempty"`

	// Pattern
	Pattern *string `json:"pattern,omitempty"`

	// Format
	Format *string `json:"format,omitempty"`

	// Items
	Items *Contractitems `json:"items,omitempty"`

	// Properties
	Properties *map[string]Contractpropertydefinition `json:"properties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contractpropertydefinition) SetField(field string, fieldValue interface{}) {
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

func (o Contractpropertydefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Contractpropertydefinition
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *[]string `json:"type,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		
		Format *string `json:"format,omitempty"`
		
		Items *Contractitems `json:"items,omitempty"`
		
		Properties *map[string]Contractpropertydefinition `json:"properties,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Pattern: o.Pattern,
		
		Format: o.Format,
		
		Items: o.Items,
		
		Properties: o.Properties,
		Alias:    (Alias)(o),
	})
}

func (o *Contractpropertydefinition) UnmarshalJSON(b []byte) error {
	var ContractpropertydefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &ContractpropertydefinitionMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ContractpropertydefinitionMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ContractpropertydefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := ContractpropertydefinitionMap["type"].([]interface{}); ok {
		VarTypeString, _ := json.Marshal(VarType)
		json.Unmarshal(VarTypeString, &o.VarType)
	}
	
	if Pattern, ok := ContractpropertydefinitionMap["pattern"].(string); ok {
		o.Pattern = &Pattern
	}
    
	if Format, ok := ContractpropertydefinitionMap["format"].(string); ok {
		o.Format = &Format
	}
    
	if Items, ok := ContractpropertydefinitionMap["items"].(map[string]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	
	if Properties, ok := ContractpropertydefinitionMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contractpropertydefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
