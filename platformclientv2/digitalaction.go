package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitalaction
type Digitalaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UpdateContactColumnActionSettings - The settings for an 'update contact column' action.
	UpdateContactColumnActionSettings *Updatecontactcolumnactionsettings `json:"updateContactColumnActionSettings,omitempty"`

	// DoNotSendActionSettings - The settings for a 'do not send' action.
	DoNotSendActionSettings *interface{} `json:"doNotSendActionSettings,omitempty"`

	// AppendToDncActionSettings - The settings for an 'Append to DNC' action.
	AppendToDncActionSettings *Appendtodncactionsettings `json:"appendToDncActionSettings,omitempty"`

	// MarkContactUncontactableActionSettings - The settings for a 'mark contact uncontactable' action.
	MarkContactUncontactableActionSettings *Markcontactuncontactableactionsettings `json:"markContactUncontactableActionSettings,omitempty"`

	// MarkContactAddressUncontactableActionSettings - The settings for an 'mark contact address uncontactable' action.
	MarkContactAddressUncontactableActionSettings *interface{} `json:"markContactAddressUncontactableActionSettings,omitempty"`

	// SetContentTemplateActionSettings - The settings for a 'Set content template' action.
	SetContentTemplateActionSettings *Setcontenttemplateactionsettings `json:"setContentTemplateActionSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Digitalaction) SetField(field string, fieldValue interface{}) {
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

func (o Digitalaction) MarshalJSON() ([]byte, error) {
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
	type Alias Digitalaction
	
	return json.Marshal(&struct { 
		UpdateContactColumnActionSettings *Updatecontactcolumnactionsettings `json:"updateContactColumnActionSettings,omitempty"`
		
		DoNotSendActionSettings *interface{} `json:"doNotSendActionSettings,omitempty"`
		
		AppendToDncActionSettings *Appendtodncactionsettings `json:"appendToDncActionSettings,omitempty"`
		
		MarkContactUncontactableActionSettings *Markcontactuncontactableactionsettings `json:"markContactUncontactableActionSettings,omitempty"`
		
		MarkContactAddressUncontactableActionSettings *interface{} `json:"markContactAddressUncontactableActionSettings,omitempty"`
		
		SetContentTemplateActionSettings *Setcontenttemplateactionsettings `json:"setContentTemplateActionSettings,omitempty"`
		Alias
	}{ 
		UpdateContactColumnActionSettings: o.UpdateContactColumnActionSettings,
		
		DoNotSendActionSettings: o.DoNotSendActionSettings,
		
		AppendToDncActionSettings: o.AppendToDncActionSettings,
		
		MarkContactUncontactableActionSettings: o.MarkContactUncontactableActionSettings,
		
		MarkContactAddressUncontactableActionSettings: o.MarkContactAddressUncontactableActionSettings,
		
		SetContentTemplateActionSettings: o.SetContentTemplateActionSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Digitalaction) UnmarshalJSON(b []byte) error {
	var DigitalactionMap map[string]interface{}
	err := json.Unmarshal(b, &DigitalactionMap)
	if err != nil {
		return err
	}
	
	if UpdateContactColumnActionSettings, ok := DigitalactionMap["updateContactColumnActionSettings"].(map[string]interface{}); ok {
		UpdateContactColumnActionSettingsString, _ := json.Marshal(UpdateContactColumnActionSettings)
		json.Unmarshal(UpdateContactColumnActionSettingsString, &o.UpdateContactColumnActionSettings)
	}
	
	if DoNotSendActionSettings, ok := DigitalactionMap["doNotSendActionSettings"].(map[string]interface{}); ok {
		DoNotSendActionSettingsString, _ := json.Marshal(DoNotSendActionSettings)
		json.Unmarshal(DoNotSendActionSettingsString, &o.DoNotSendActionSettings)
	}
	
	if AppendToDncActionSettings, ok := DigitalactionMap["appendToDncActionSettings"].(map[string]interface{}); ok {
		AppendToDncActionSettingsString, _ := json.Marshal(AppendToDncActionSettings)
		json.Unmarshal(AppendToDncActionSettingsString, &o.AppendToDncActionSettings)
	}
	
	if MarkContactUncontactableActionSettings, ok := DigitalactionMap["markContactUncontactableActionSettings"].(map[string]interface{}); ok {
		MarkContactUncontactableActionSettingsString, _ := json.Marshal(MarkContactUncontactableActionSettings)
		json.Unmarshal(MarkContactUncontactableActionSettingsString, &o.MarkContactUncontactableActionSettings)
	}
	
	if MarkContactAddressUncontactableActionSettings, ok := DigitalactionMap["markContactAddressUncontactableActionSettings"].(map[string]interface{}); ok {
		MarkContactAddressUncontactableActionSettingsString, _ := json.Marshal(MarkContactAddressUncontactableActionSettings)
		json.Unmarshal(MarkContactAddressUncontactableActionSettingsString, &o.MarkContactAddressUncontactableActionSettings)
	}
	
	if SetContentTemplateActionSettings, ok := DigitalactionMap["setContentTemplateActionSettings"].(map[string]interface{}); ok {
		SetContentTemplateActionSettingsString, _ := json.Marshal(SetContentTemplateActionSettings)
		json.Unmarshal(SetContentTemplateActionSettingsString, &o.SetContentTemplateActionSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Digitalaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
