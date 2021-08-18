package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsnetworktypeip
type Trunkmetricsnetworktypeip struct { 
	// Address - Assigned IP Address for the interface
	Address *string `json:"address,omitempty"`


	// ErrorInfo - Information about the error.
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

func (u *Trunkmetricsnetworktypeip) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsnetworktypeip

	

	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
