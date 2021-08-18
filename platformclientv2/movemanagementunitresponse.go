package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Movemanagementunitresponse
type Movemanagementunitresponse struct { 
	// BusinessUnit - The new business unit
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// Status - The status of the move.  Will always be 'Processing' unless the Management Unit is already in the requested Business Unit in which case it will be 'Complete'
	Status *string `json:"status,omitempty"`

}

func (u *Movemanagementunitresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Movemanagementunitresponse

	

	return json.Marshal(&struct { 
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		BusinessUnit: u.BusinessUnit,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Movemanagementunitresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
