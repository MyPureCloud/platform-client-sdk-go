package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Embeddedintegration
type Embeddedintegration struct { 
	// EnableWhitelist
	EnableWhitelist *bool `json:"enableWhitelist,omitempty"`


	// DomainWhitelist
	DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`

}

func (u *Embeddedintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Embeddedintegration

	

	return json.Marshal(&struct { 
		EnableWhitelist *bool `json:"enableWhitelist,omitempty"`
		
		DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`
		*Alias
	}{ 
		EnableWhitelist: u.EnableWhitelist,
		
		DomainWhitelist: u.DomainWhitelist,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Embeddedintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
