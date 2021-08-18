package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Movemanagementunitrequest
type Movemanagementunitrequest struct { 
	// BusinessUnitId - The ID of the business unit to which to move the management unit
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

func (u *Movemanagementunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Movemanagementunitrequest

	

	return json.Marshal(&struct { 
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		*Alias
	}{ 
		BusinessUnitId: u.BusinessUnitId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Movemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
