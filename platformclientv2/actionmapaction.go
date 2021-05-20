package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Actionmapaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
