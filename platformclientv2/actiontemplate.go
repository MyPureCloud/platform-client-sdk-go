package platformclientv2
import (
	"time"
	"encoding/json"
)

// Actiontemplate
type Actiontemplate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name of the action template.
	Name *string `json:"name,omitempty"`


	// Description - Description of the action template's functionality.
	Description *string `json:"description,omitempty"`


	// MediaType - Media type of action described by the action template.
	MediaType *string `json:"mediaType,omitempty"`


	// State - Whether the action template is currently active, inactive or deleted.
	State *string `json:"state,omitempty"`


	// ContentOffer - Properties used to configure an action of type content offer
	ContentOffer *Contentoffer `json:"contentOffer,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Date when action template was created in ISO-8601 format.
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date when action template was last modified in ISO-8601 format.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actiontemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
