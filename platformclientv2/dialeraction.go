package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeraction
type Dialeraction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of this DialerAction.
	VarType *string `json:"type,omitempty"`

	// ActionTypeName - Additional type specification for this DialerAction.
	ActionTypeName *string `json:"actionTypeName,omitempty"`

	// UpdateOption - Specifies how a contact attribute should be updated. Required for MODIFY_CONTACT_ATTRIBUTE.
	UpdateOption *string `json:"updateOption,omitempty"`

	// Properties - A map of key-value pairs pertinent to the DialerAction. Different types of DialerActions require different properties. MODIFY_CONTACT_ATTRIBUTE with an updateOption of SET takes a contact column as the key and accepts any value. SCHEDULE_CALLBACK takes a key 'callbackOffset' that specifies how far in the future the callback should be scheduled, in minutes. SET_CALLER_ID takes two keys: 'callerAddress', which should be the caller id phone number, and 'callerName'. For either key, you can also specify a column on the contact to get the value from. To do this, specify 'contact.Column', where 'Column' is the name of the contact column from which to get the value. SET_SKILLS takes a key 'skills' with an array of skill ids wrapped into a string (Example: {'skills': '['skillIdHere']'} ).
	Properties *map[string]string `json:"properties,omitempty"`

	// DataAction - The Data Action to use for this action. Required for a dataActionBehavior.
	DataAction *Domainentityref `json:"dataAction,omitempty"`

	// ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionBehavior.
	ContactColumnToDataActionFieldMappings *[]Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`

	// ContactIdField - The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionBehavior.
	ContactIdField *string `json:"contactIdField,omitempty"`

	// CallAnalysisResultField - The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionBehavior.
	CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`

	// AgentWrapupField - The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionBehavior.
	AgentWrapupField *string `json:"agentWrapupField,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialeraction) SetField(field string, fieldValue interface{}) {
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

func (o Dialeraction) MarshalJSON() ([]byte, error) {
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
	type Alias Dialeraction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ActionTypeName *string `json:"actionTypeName,omitempty"`
		
		UpdateOption *string `json:"updateOption,omitempty"`
		
		Properties *map[string]string `json:"properties,omitempty"`
		
		DataAction *Domainentityref `json:"dataAction,omitempty"`
		
		ContactColumnToDataActionFieldMappings *[]Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
		
		ContactIdField *string `json:"contactIdField,omitempty"`
		
		CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`
		
		AgentWrapupField *string `json:"agentWrapupField,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		ActionTypeName: o.ActionTypeName,
		
		UpdateOption: o.UpdateOption,
		
		Properties: o.Properties,
		
		DataAction: o.DataAction,
		
		ContactColumnToDataActionFieldMappings: o.ContactColumnToDataActionFieldMappings,
		
		ContactIdField: o.ContactIdField,
		
		CallAnalysisResultField: o.CallAnalysisResultField,
		
		AgentWrapupField: o.AgentWrapupField,
		Alias:    (Alias)(o),
	})
}

func (o *Dialeraction) UnmarshalJSON(b []byte) error {
	var DialeractionMap map[string]interface{}
	err := json.Unmarshal(b, &DialeractionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DialeractionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ActionTypeName, ok := DialeractionMap["actionTypeName"].(string); ok {
		o.ActionTypeName = &ActionTypeName
	}
    
	if UpdateOption, ok := DialeractionMap["updateOption"].(string); ok {
		o.UpdateOption = &UpdateOption
	}
    
	if Properties, ok := DialeractionMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if DataAction, ok := DialeractionMap["dataAction"].(map[string]interface{}); ok {
		DataActionString, _ := json.Marshal(DataAction)
		json.Unmarshal(DataActionString, &o.DataAction)
	}
	
	if ContactColumnToDataActionFieldMappings, ok := DialeractionMap["contactColumnToDataActionFieldMappings"].([]interface{}); ok {
		ContactColumnToDataActionFieldMappingsString, _ := json.Marshal(ContactColumnToDataActionFieldMappings)
		json.Unmarshal(ContactColumnToDataActionFieldMappingsString, &o.ContactColumnToDataActionFieldMappings)
	}
	
	if ContactIdField, ok := DialeractionMap["contactIdField"].(string); ok {
		o.ContactIdField = &ContactIdField
	}
    
	if CallAnalysisResultField, ok := DialeractionMap["callAnalysisResultField"].(string); ok {
		o.CallAnalysisResultField = &CallAnalysisResultField
	}
    
	if AgentWrapupField, ok := DialeractionMap["agentWrapupField"].(string); ok {
		o.AgentWrapupField = &AgentWrapupField
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeraction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
