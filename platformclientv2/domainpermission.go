package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainpermission
type Domainpermission struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Domain
	Domain *string `json:"domain,omitempty"`

	// EntityType
	EntityType *string `json:"entityType,omitempty"`

	// Action
	Action *string `json:"action,omitempty"`

	// Label
	Label *string `json:"label,omitempty"`

	// AllowsConditions
	AllowsConditions *bool `json:"allowsConditions,omitempty"`

	// DivisionAware
	DivisionAware *bool `json:"divisionAware,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainpermission) SetField(field string, fieldValue interface{}) {
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

func (o Domainpermission) MarshalJSON() ([]byte, error) {
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
	type Alias Domainpermission
	
	return json.Marshal(&struct { 
		Domain *string `json:"domain,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		AllowsConditions *bool `json:"allowsConditions,omitempty"`
		
		DivisionAware *bool `json:"divisionAware,omitempty"`
		Alias
	}{ 
		Domain: o.Domain,
		
		EntityType: o.EntityType,
		
		Action: o.Action,
		
		Label: o.Label,
		
		AllowsConditions: o.AllowsConditions,
		
		DivisionAware: o.DivisionAware,
		Alias:    (Alias)(o),
	})
}

func (o *Domainpermission) UnmarshalJSON(b []byte) error {
	var DomainpermissionMap map[string]interface{}
	err := json.Unmarshal(b, &DomainpermissionMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := DomainpermissionMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if EntityType, ok := DomainpermissionMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if Action, ok := DomainpermissionMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Label, ok := DomainpermissionMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if AllowsConditions, ok := DomainpermissionMap["allowsConditions"].(bool); ok {
		o.AllowsConditions = &AllowsConditions
	}
    
	if DivisionAware, ok := DomainpermissionMap["divisionAware"].(bool); ok {
		o.DivisionAware = &DivisionAware
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainpermission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
