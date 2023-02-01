package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapaction
type Actionmapaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`

	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`

	// ActionTargetId - Action target ID.
	ActionTargetId *string `json:"actionTargetId,omitempty"`

	// IsPacingEnabled - Whether this action should be throttled.
	IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`

	// Props - Additional properties.
	Props *Actionproperties `json:"props,omitempty"`

	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`

	// WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
	WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`

	// OpenActionFields - Admin-configurable fields of an open action.
	OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionmapaction) SetField(field string, fieldValue interface{}) {
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

func (o Actionmapaction) MarshalJSON() ([]byte, error) {
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
	type Alias Actionmapaction
	
	return json.Marshal(&struct { 
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ActionTargetId *string `json:"actionTargetId,omitempty"`
		
		IsPacingEnabled *bool `json:"isPacingEnabled,omitempty"`
		
		Props *Actionproperties `json:"props,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		
		OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
		Alias
	}{ 
		ActionTemplate: o.ActionTemplate,
		
		MediaType: o.MediaType,
		
		ActionTargetId: o.ActionTargetId,
		
		IsPacingEnabled: o.IsPacingEnabled,
		
		Props: o.Props,
		
		ArchitectFlowFields: o.ArchitectFlowFields,
		
		WebMessagingOfferFields: o.WebMessagingOfferFields,
		
		OpenActionFields: o.OpenActionFields,
		Alias:    (Alias)(o),
	})
}

func (o *Actionmapaction) UnmarshalJSON(b []byte) error {
	var ActionmapactionMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapactionMap)
	if err != nil {
		return err
	}
	
	if ActionTemplate, ok := ActionmapactionMap["actionTemplate"].(map[string]interface{}); ok {
		ActionTemplateString, _ := json.Marshal(ActionTemplate)
		json.Unmarshal(ActionTemplateString, &o.ActionTemplate)
	}
	
	if MediaType, ok := ActionmapactionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ActionTargetId, ok := ActionmapactionMap["actionTargetId"].(string); ok {
		o.ActionTargetId = &ActionTargetId
	}
    
	if IsPacingEnabled, ok := ActionmapactionMap["isPacingEnabled"].(bool); ok {
		o.IsPacingEnabled = &IsPacingEnabled
	}
    
	if Props, ok := ActionmapactionMap["props"].(map[string]interface{}); ok {
		PropsString, _ := json.Marshal(Props)
		json.Unmarshal(PropsString, &o.Props)
	}
	
	if ArchitectFlowFields, ok := ActionmapactionMap["architectFlowFields"].(map[string]interface{}); ok {
		ArchitectFlowFieldsString, _ := json.Marshal(ArchitectFlowFields)
		json.Unmarshal(ArchitectFlowFieldsString, &o.ArchitectFlowFields)
	}
	
	if WebMessagingOfferFields, ok := ActionmapactionMap["webMessagingOfferFields"].(map[string]interface{}); ok {
		WebMessagingOfferFieldsString, _ := json.Marshal(WebMessagingOfferFields)
		json.Unmarshal(WebMessagingOfferFieldsString, &o.WebMessagingOfferFields)
	}
	
	if OpenActionFields, ok := ActionmapactionMap["openActionFields"].(map[string]interface{}); ok {
		OpenActionFieldsString, _ := json.Marshal(OpenActionFields)
		json.Unmarshal(OpenActionFieldsString, &o.OpenActionFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
