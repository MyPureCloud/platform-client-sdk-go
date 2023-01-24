package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataactionconditionsettings
type Dataactionconditionsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DataActionId - The Data Action Id to use for this condition.
	DataActionId *string `json:"dataActionId,omitempty"`

	// ContactIdField - The input field from the data action that the contactId will be passed into.
	ContactIdField *string `json:"contactIdField,omitempty"`

	// DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data.
	DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`

	// Predicates - A list of predicates defining the comparisons to use for this condition.
	Predicates *[]Digitaldataactionconditionpredicate `json:"predicates,omitempty"`

	// ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields.
	ContactColumnToDataActionFieldMappings *[]Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dataactionconditionsettings) SetField(field string, fieldValue interface{}) {
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

func (o Dataactionconditionsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Dataactionconditionsettings
	
	return json.Marshal(&struct { 
		DataActionId *string `json:"dataActionId,omitempty"`
		
		ContactIdField *string `json:"contactIdField,omitempty"`
		
		DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`
		
		Predicates *[]Digitaldataactionconditionpredicate `json:"predicates,omitempty"`
		
		ContactColumnToDataActionFieldMappings *[]Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
		Alias
	}{ 
		DataActionId: o.DataActionId,
		
		ContactIdField: o.ContactIdField,
		
		DataNotFoundResolution: o.DataNotFoundResolution,
		
		Predicates: o.Predicates,
		
		ContactColumnToDataActionFieldMappings: o.ContactColumnToDataActionFieldMappings,
		Alias:    (Alias)(o),
	})
}

func (o *Dataactionconditionsettings) UnmarshalJSON(b []byte) error {
	var DataactionconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DataactionconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if DataActionId, ok := DataactionconditionsettingsMap["dataActionId"].(string); ok {
		o.DataActionId = &DataActionId
	}
    
	if ContactIdField, ok := DataactionconditionsettingsMap["contactIdField"].(string); ok {
		o.ContactIdField = &ContactIdField
	}
    
	if DataNotFoundResolution, ok := DataactionconditionsettingsMap["dataNotFoundResolution"].(bool); ok {
		o.DataNotFoundResolution = &DataNotFoundResolution
	}
    
	if Predicates, ok := DataactionconditionsettingsMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	
	if ContactColumnToDataActionFieldMappings, ok := DataactionconditionsettingsMap["contactColumnToDataActionFieldMappings"].([]interface{}); ok {
		ContactColumnToDataActionFieldMappingsString, _ := json.Marshal(ContactColumnToDataActionFieldMappings)
		json.Unmarshal(ContactColumnToDataActionFieldMappingsString, &o.ContactColumnToDataActionFieldMappings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dataactionconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
