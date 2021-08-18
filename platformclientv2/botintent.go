package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botintent - A botConnector's bot intention
type Botintent struct { 
	// Name - The name of this intent.  This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
	Name *string `json:"name,omitempty"`


	// Slots - Optional returned data values associated with this intent, limit of 50.
	Slots *map[string]Botslot `json:"slots,omitempty"`

}

func (u *Botintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botintent

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Slots *map[string]Botslot `json:"slots,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Slots: u.Slots,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Botintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
