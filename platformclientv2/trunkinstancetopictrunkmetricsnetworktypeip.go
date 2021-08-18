package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkmetricsnetworktypeip
type Trunkinstancetopictrunkmetricsnetworktypeip struct { 
	// Address
	Address *string `json:"address,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`

}

func (u *Trunkinstancetopictrunkmetricsnetworktypeip) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkmetricsnetworktypeip

	

	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
