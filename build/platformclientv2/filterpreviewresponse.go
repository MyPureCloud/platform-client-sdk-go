package platformclientv2
import (
	"encoding/json"
)

// Filterpreviewresponse
type Filterpreviewresponse struct { 
	// FilteredContacts
	FilteredContacts *int `json:"filteredContacts,omitempty"`


	// TotalContacts
	TotalContacts *int `json:"totalContacts,omitempty"`


	// Preview
	Preview *[]Dialercontact `json:"preview,omitempty"`

}

// String returns a JSON representation of the model
func (o *Filterpreviewresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
