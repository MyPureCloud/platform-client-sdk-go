package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitalaction
type Digitalaction struct { 
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

func (o *Digitalaction) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		UpdateContactColumnActionSettings: o.UpdateContactColumnActionSettings,
		
		DoNotSendActionSettings: o.DoNotSendActionSettings,
		
		AppendToDncActionSettings: o.AppendToDncActionSettings,
		
		MarkContactUncontactableActionSettings: o.MarkContactUncontactableActionSettings,
		
		MarkContactAddressUncontactableActionSettings: o.MarkContactAddressUncontactableActionSettings,
		
		SetContentTemplateActionSettings: o.SetContentTemplateActionSettings,
		Alias:    (*Alias)(o),
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
