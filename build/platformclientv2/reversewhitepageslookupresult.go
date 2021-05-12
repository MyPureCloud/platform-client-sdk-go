package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Reversewhitepageslookupresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
