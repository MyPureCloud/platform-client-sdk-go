package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Filterpreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Filterpreviewresponse

	

	return json.Marshal(&struct { 
		FilteredContacts *int `json:"filteredContacts,omitempty"`
		
		TotalContacts *int `json:"totalContacts,omitempty"`
		
		Preview *[]Dialercontact `json:"preview,omitempty"`
		*Alias
	}{ 
		FilteredContacts: u.FilteredContacts,
		
		TotalContacts: u.TotalContacts,
		
		Preview: u.Preview,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Filterpreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
