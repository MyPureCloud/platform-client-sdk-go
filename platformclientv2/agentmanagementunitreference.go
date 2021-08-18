package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentmanagementunitreference - A reference from agent to management unit
type Agentmanagementunitreference struct { 
	// User - The user (agent) for whom the management unit was requested
	User *Userreference `json:"user,omitempty"`


	// ManagementUnit - The management to which the user (agent) belongs
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// BusinessUnit - The business unit to which the user (agent) belongs. Populate with expand=businessUnit
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`

}

func (u *Agentmanagementunitreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentmanagementunitreference

	

	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		ManagementUnit: u.ManagementUnit,
		
		BusinessUnit: u.BusinessUnit,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentmanagementunitreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
