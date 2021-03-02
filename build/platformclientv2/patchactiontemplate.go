package platformclientv2
import (
	"encoding/json"
)

// Patchactiontemplate
type Patchactiontemplate struct { 
	// Name - Name of the action template.
	Name *string `json:"name,omitempty"`


	// Description - Description of the action template's functionality.
	Description *string `json:"description,omitempty"`


	// MediaType - Media type of action described by the action template.
	MediaType *string `json:"mediaType,omitempty"`


	// State - Whether the action template is currently active, inactive or deleted.
	State *string `json:"state,omitempty"`


	// ContentOffer - Properties used to configure an action of type content offer
	ContentOffer *Patchcontentoffer `json:"contentOffer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchactiontemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
