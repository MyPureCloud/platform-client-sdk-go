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

}

func (u *Actionmapaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapaction

	

	return json.Marshal(&struct { 
		ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`
		
		WebMessagingOfferFields *Webmessagingofferfields `json:"webMessagingOfferFields,omitempty"`
		*Alias
	}{ 
		ActionTemplate: u.ActionTemplate,
		
		MediaType: u.MediaType,
		
		ArchitectFlowFields: u.ArchitectFlowFields,
		
		WebMessagingOfferFields: u.WebMessagingOfferFields,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actionmapaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
