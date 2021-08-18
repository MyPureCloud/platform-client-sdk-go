package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonesreboot
type Phonesreboot struct { 
	// PhoneIds - The list of phone Ids to reboot.
	PhoneIds *[]string `json:"phoneIds,omitempty"`


	// SiteId - ID of the site for which to reboot all phones at that site. no.active.edge and phone.cannot.resolve errors are ignored.
	SiteId *string `json:"siteId,omitempty"`

}

func (u *Phonesreboot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonesreboot

	

	return json.Marshal(&struct { 
		PhoneIds *[]string `json:"phoneIds,omitempty"`
		
		SiteId *string `json:"siteId,omitempty"`
		*Alias
	}{ 
		PhoneIds: u.PhoneIds,
		
		SiteId: u.SiteId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonesreboot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
