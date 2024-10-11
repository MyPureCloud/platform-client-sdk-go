package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Intentdefinition
type Intentdefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the intent.
	Id *string `json:"id,omitempty"`

	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`

	// Description - The description of the intent.
	Description *string `json:"description,omitempty"`

	// EntityTypeBindings - The bindings for the named entity types used in this intent.This field is mutually exclusive with entityNameReferences and entities
	EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`

	// EntityNameReferences - The references for the named entity used in this intent.This field is mutually exclusive with entityTypeBindings
	EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`

	// Utterances - The utterances that act as training phrases for the intent.
	Utterances *[]Nluutterance `json:"utterances,omitempty"`

	// AdditionalLanguages - Additional languages for intents
	AdditionalLanguages *map[string]Additionallanguagesintent `json:"additionalLanguages,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Intentdefinition) SetField(field string, fieldValue interface{}) {
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

func (o Intentdefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Intentdefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`
		
		EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`
		
		Utterances *[]Nluutterance `json:"utterances,omitempty"`
		
		AdditionalLanguages *map[string]Additionallanguagesintent `json:"additionalLanguages,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		EntityTypeBindings: o.EntityTypeBindings,
		
		EntityNameReferences: o.EntityNameReferences,
		
		Utterances: o.Utterances,
		
		AdditionalLanguages: o.AdditionalLanguages,
		Alias:    (Alias)(o),
	})
}

func (o *Intentdefinition) UnmarshalJSON(b []byte) error {
	var IntentdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &IntentdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := IntentdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := IntentdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := IntentdefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if EntityTypeBindings, ok := IntentdefinitionMap["entityTypeBindings"].([]interface{}); ok {
		EntityTypeBindingsString, _ := json.Marshal(EntityTypeBindings)
		json.Unmarshal(EntityTypeBindingsString, &o.EntityTypeBindings)
	}
	
	if EntityNameReferences, ok := IntentdefinitionMap["entityNameReferences"].([]interface{}); ok {
		EntityNameReferencesString, _ := json.Marshal(EntityNameReferences)
		json.Unmarshal(EntityNameReferencesString, &o.EntityNameReferences)
	}
	
	if Utterances, ok := IntentdefinitionMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	
	if AdditionalLanguages, ok := IntentdefinitionMap["additionalLanguages"].(map[string]interface{}); ok {
		AdditionalLanguagesString, _ := json.Marshal(AdditionalLanguages)
		json.Unmarshal(AdditionalLanguagesString, &o.AdditionalLanguages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Intentdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
