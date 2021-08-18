package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappdefinition - A WhatsApp messaging template definition as defined in the WhatsApp Business Manager
type Whatsappdefinition struct { 
	// Name - The messaging template name.
	Name *string `json:"name,omitempty"`


	// Namespace - The messaging template namespace.
	Namespace *string `json:"namespace,omitempty"`


	// Language - The messaging template language configured for this template. This is a WhatsApp specific value. For example, 'en_US'
	Language *string `json:"language,omitempty"`

}

func (u *Whatsappdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappdefinition

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		Language *string `json:"language,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Namespace: u.Namespace,
		
		Language: u.Language,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Whatsappdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
