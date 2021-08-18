package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluinfo
type Nluinfo struct { 
	// Domain
	Domain *Addressableentityref `json:"domain,omitempty"`


	// Version
	Version **Nludomainversion `json:"version,omitempty"`


	// Intents
	Intents *[]Intent `json:"intents,omitempty"`

}

func (u *Nluinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluinfo

	

	return json.Marshal(&struct { 
		Domain *Addressableentityref `json:"domain,omitempty"`
		
		Version **Nludomainversion `json:"version,omitempty"`
		
		Intents *[]Intent `json:"intents,omitempty"`
		*Alias
	}{ 
		Domain: u.Domain,
		
		Version: u.Version,
		
		Intents: u.Intents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
