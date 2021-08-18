package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Patchactiontemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactiontemplate

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ContentOffer *Patchcontentoffer `json:"contentOffer,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		MediaType: u.MediaType,
		
		State: u.State,
		
		ContentOffer: u.ContentOffer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchactiontemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
