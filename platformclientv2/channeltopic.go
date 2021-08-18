package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Channeltopic
type Channeltopic struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Channeltopic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Channeltopic

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Channeltopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
