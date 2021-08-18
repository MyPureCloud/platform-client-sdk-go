package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportedcontentreference - Reference to supported content profile associated with the integration
type Supportedcontentreference struct { 
	// Id - The SupportedContent unique identifier associated with this integration
	Id *string `json:"id,omitempty"`


	// Name - The SupportedContent profile name
	Name *string `json:"name,omitempty"`


	// SelfUri - The SupportedContent profile URI
	SelfUri *string `json:"selfUri,omitempty"`


	// MediaTypes - Media types definition for the supported content
	MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`

}

func (u *Supportedcontentreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportedcontentreference

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SelfUri: u.SelfUri,
		
		MediaTypes: u.MediaTypes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Supportedcontentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
