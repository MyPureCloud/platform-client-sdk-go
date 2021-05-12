package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentscomplete) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
