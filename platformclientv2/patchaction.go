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

func (u *Patchaction) MarshalJSON() ([]byte, error) {
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
		MediaType: u.MediaType,
		
		ActionTemplate: u.ActionTemplate,
		
		ArchitectFlowFields: u.ArchitectFlowFields,
		
		WebMessagingOfferFields: u.WebMessagingOfferFields,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
