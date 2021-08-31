package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchaction
type Patchaction struct { 
	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`


	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`


	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`


	// WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
	WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`

}

func (o *Patchaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchaction
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		*Alias
	}{ 
		MediaType: o.MediaType,
		
		ActionTemplate: o.ActionTemplate,
		
		ArchitectFlowFields: o.ArchitectFlowFields,
		
		WebMessagingOfferFields: o.WebMessagingOfferFields,
		Alias:    (*Alias)(o),
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
	
	if ArchitectFlowFields, ok := PatchactionMap["architectFlowFields"].(map[string]interface{}); ok {
		ArchitectFlowFieldsString, _ := json.Marshal(ArchitectFlowFields)
		json.Unmarshal(ArchitectFlowFieldsString, &o.ArchitectFlowFields)
	}
	
	if WebMessagingOfferFields, ok := PatchactionMap["webMessagingOfferFields"].(map[string]interface{}); ok {
		WebMessagingOfferFieldsString, _ := json.Marshal(WebMessagingOfferFields)
		json.Unmarshal(WebMessagingOfferFieldsString, &o.WebMessagingOfferFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
