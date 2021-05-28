package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
