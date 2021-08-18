package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceeventorganizationpresence
type Presenceeventorganizationpresence struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

}

func (u *Presenceeventorganizationpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presenceeventorganizationpresence

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		SystemPresence: u.SystemPresence,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Presenceeventorganizationpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
