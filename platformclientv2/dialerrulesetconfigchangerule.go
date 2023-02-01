package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangerule
type Dialerrulesetconfigchangerule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Conditions - The list of rule conditions; all must evaluate to true to trigger the rule actions
	Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// Id - The identifier of the rule
	Id *string `json:"id,omitempty"`

	// Name - The name of the rule
	Name *string `json:"name,omitempty"`

	// Order - The ranked order of the rule; rules are processed from lowest number to highest
	Order *int `json:"order,omitempty"`

	// Category - The category of the rule
	Category *string `json:"category,omitempty"`

	// Actions - The list of rule actions to be taken if the conditions are true
	Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialerrulesetconfigchangerule) SetField(field string, fieldValue interface{}) {
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

func (o Dialerrulesetconfigchangerule) MarshalJSON() ([]byte, error) {
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
	type Alias Dialerrulesetconfigchangerule
	
	return json.Marshal(&struct { 
		Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`
		Alias
	}{ 
		Conditions: o.Conditions,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		Order: o.Order,
		
		Category: o.Category,
		
		Actions: o.Actions,
		Alias:    (Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangerule) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeruleMap)
	if err != nil {
		return err
	}
	
	if Conditions, ok := DialerrulesetconfigchangeruleMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if AdditionalProperties, ok := DialerrulesetconfigchangeruleMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialerrulesetconfigchangeruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerrulesetconfigchangeruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Order, ok := DialerrulesetconfigchangeruleMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if Category, ok := DialerrulesetconfigchangeruleMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Actions, ok := DialerrulesetconfigchangeruleMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
