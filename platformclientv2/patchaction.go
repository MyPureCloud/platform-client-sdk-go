package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchaction
type Patchaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`

	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`

	// ActionTargetId - Action target ID.
	ActionTargetId *string `json:"actionTargetId,omitempty"`

	// IsPacingEnabled - Whether this action should be throttled.
	IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`

	// Props - Additional properties.
	Props *Patchactionproperties `json:"props,omitempty"`

	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`

	// WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
	WebMessagingOfferFields *Patchwebmessagingofferfields `json:"webMessagingOfferFields,omitempty"`

	// OpenActionFields - Admin-configurable fields of an open action.
	OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchaction) SetField(field string, fieldValue interface{}) {
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

func (o Patchaction) MarshalJSON() ([]byte, error) {
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
	type Alias Patchaction
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		ActionTargetId *string `json:"actionTargetId,omitempty"`
		
		IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`
		
		Props *Patchactionproperties `json:"props,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Patchwebmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		
		OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
		Alias
	}{ 
		MediaType: o.MediaType,
		
		ActionTemplate: o.ActionTemplate,
		
		ActionTargetId: o.ActionTargetId,
		
		IsPacingEnabled: o.IsPacingEnabled,
		
		Props: o.Props,
		
		ArchitectFlowFields: o.ArchitectFlowFields,
		
		WebMessagingOfferFields: o.WebMessagingOfferFields,
		
		OpenActionFields: o.OpenActionFields,
		Alias:    (Alias)(o),
	})
}

func (o *Patchaction) UnmarshalJSON(b []byte) error {
	var PatchactionMap map[string]interface{}
	err := json.Unmarshal(b, &PatchactionMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := PatchactionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ActionTemplate, ok := PatchactionMap["actionTemplate"].(map[string]interface{}); ok {
		ActionTemplateString, _ := json.Marshal(ActionTemplate)
		json.Unmarshal(ActionTemplateString, &o.ActionTemplate)
	}
	
	if ActionTargetId, ok := PatchactionMap["actionTargetId"].(string); ok {
		o.ActionTargetId = &ActionTargetId
	}
    
	if IsPacingEnabled, ok := PatchactionMap["isPacingEnabled"].(bool); ok {
		o.IsPacingEnabled = &IsPacingEnabled
	}
    
	if Props, ok := PatchactionMap["props"].(map[string]interface{}); ok {
		PropsString, _ := json.Marshal(Props)
		json.Unmarshal(PropsString, &o.Props)
	}
	
	if ArchitectFlowFields, ok := PatchactionMap["architectFlowFields"].(map[string]interface{}); ok {
		ArchitectFlowFieldsString, _ := json.Marshal(ArchitectFlowFields)
		json.Unmarshal(ArchitectFlowFieldsString, &o.ArchitectFlowFields)
	}
	
	if WebMessagingOfferFields, ok := PatchactionMap["webMessagingOfferFields"].(map[string]interface{}); ok {
		WebMessagingOfferFieldsString, _ := json.Marshal(WebMessagingOfferFields)
		json.Unmarshal(WebMessagingOfferFieldsString, &o.WebMessagingOfferFields)
	}
	
	if OpenActionFields, ok := PatchactionMap["openActionFields"].(map[string]interface{}); ok {
		OpenActionFieldsString, _ := json.Marshal(OpenActionFields)
		json.Unmarshal(OpenActionFieldsString, &o.OpenActionFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
