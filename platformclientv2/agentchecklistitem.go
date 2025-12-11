package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentchecklistitem
type Agentchecklistitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the checklist item.
	Id *string `json:"id,omitempty"`

	// Name - Name of the checklist item.
	Name *string `json:"name,omitempty"`

	// Description - Description of the checklist item.
	Description *string `json:"description,omitempty"`

	// AutomatedCheckEnabled - Flag to indicate whether automated check is enabled for this checklist item.
	AutomatedCheckEnabled *bool `json:"automatedCheckEnabled,omitempty"`

	// Important - Flag to indicate whether this checklist item is marked as important.
	Important *bool `json:"important,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentchecklistitem) SetField(field string, fieldValue interface{}) {
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

func (o Agentchecklistitem) MarshalJSON() ([]byte, error) {
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
	type Alias Agentchecklistitem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		AutomatedCheckEnabled *bool `json:"automatedCheckEnabled,omitempty"`
		
		Important *bool `json:"important,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		AutomatedCheckEnabled: o.AutomatedCheckEnabled,
		
		Important: o.Important,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentchecklistitem) UnmarshalJSON(b []byte) error {
	var AgentchecklistitemMap map[string]interface{}
	err := json.Unmarshal(b, &AgentchecklistitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentchecklistitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentchecklistitemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := AgentchecklistitemMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if AutomatedCheckEnabled, ok := AgentchecklistitemMap["automatedCheckEnabled"].(bool); ok {
		o.AutomatedCheckEnabled = &AutomatedCheckEnabled
	}
    
	if Important, ok := AgentchecklistitemMap["important"].(bool); ok {
		o.Important = &Important
	}
    
	if SelfUri, ok := AgentchecklistitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentchecklistitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
