package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicwfmmoveagentscomplete
type Wfmmoveagentscompletetopicwfmmoveagentscomplete struct { 
	// RequestingUser
	RequestingUser *Wfmmoveagentscompletetopicuserreference `json:"requestingUser,omitempty"`


	// DestinationManagementUnit
	DestinationManagementUnit *Wfmmoveagentscompletetopicmanagementunit `json:"destinationManagementUnit,omitempty"`


	// Results
	Results *[]Wfmmoveagentscompletetopicwfmmoveagentdata `json:"results,omitempty"`

}

func (u *Wfmmoveagentscompletetopicwfmmoveagentscomplete) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmoveagentscompletetopicwfmmoveagentscomplete

	

	return json.Marshal(&struct { 
		RequestingUser *Wfmmoveagentscompletetopicuserreference `json:"requestingUser,omitempty"`
		
		DestinationManagementUnit *Wfmmoveagentscompletetopicmanagementunit `json:"destinationManagementUnit,omitempty"`
		
		Results *[]Wfmmoveagentscompletetopicwfmmoveagentdata `json:"results,omitempty"`
		*Alias
	}{ 
		RequestingUser: u.RequestingUser,
		
		DestinationManagementUnit: u.DestinationManagementUnit,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentscomplete) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
