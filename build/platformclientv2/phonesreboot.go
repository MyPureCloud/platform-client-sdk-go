package platformclientv2
import (
	"encoding/json"
)

// Phonesreboot
type Phonesreboot struct { 
	// PhoneIds - The list of phone Ids to reboot.
	PhoneIds *[]string `json:"phoneIds,omitempty"`


	// SiteId - ID of the site for which to reboot all phones at that site. no.active.edge and phone.cannot.resolve errors are ignored.
	SiteId *string `json:"siteId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonesreboot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
