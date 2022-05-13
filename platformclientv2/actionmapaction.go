package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapaction
type Actionmapaction struct { 
	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`


	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`


	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`


	// WebMessagingOfferFields - Admin-configurable fields of a web messaging offer action.
	WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`


	// OpenActionFields - Admin-configurable fields of an open action.
	OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`

}

func (o *Actionmapaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapaction
	
	return json.Marshal(&struct { 
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		
		OpenActionFields *Openactionfields `json:"openActionFields,omitempty"`
		*Alias
	}{ 
		ActionTemplate: o.ActionTemplate,
		
		MediaType: o.MediaType,
		
		ArchitectFlowFields: o.ArchitectFlowFields,
		
		WebMessagingOfferFields: o.WebMessagingOfferFields,
		
		OpenActionFields: o.OpenActionFields,
		Alias:    (*Alias)(o),
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
