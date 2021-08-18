package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reversewhitepageslookupresult
type Reversewhitepageslookupresult struct { 
	// Contacts
	Contacts *[]Externalcontact `json:"contacts,omitempty"`


	// ExternalOrganizations
	ExternalOrganizations *[]Externalorganization `json:"externalOrganizations,omitempty"`

}

func (u *Reversewhitepageslookupresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reversewhitepageslookupresult

	

	return json.Marshal(&struct { 
		Contacts *[]Externalcontact `json:"contacts,omitempty"`
		
		ExternalOrganizations *[]Externalorganization `json:"externalOrganizations,omitempty"`
		*Alias
	}{ 
		Contacts: u.Contacts,
		
		ExternalOrganizations: u.ExternalOrganizations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reversewhitepageslookupresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
