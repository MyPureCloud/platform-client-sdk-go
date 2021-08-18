package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkshifttradestateupdaterequest
type Bulkshifttradestateupdaterequest struct { 
	// Entities - The shift trades to update
	Entities *[]Bulkupdateshifttradestaterequestitem `json:"entities,omitempty"`

}

func (u *Bulkshifttradestateupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkshifttradestateupdaterequest

	

	return json.Marshal(&struct { 
		Entities *[]Bulkupdateshifttradestaterequestitem `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bulkshifttradestateupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
