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

func (o *Agentmanagementunitreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentmanagementunitreference
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentmanagementunitreference) UnmarshalJSON(b []byte) error {
	var AgentmanagementunitreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &AgentmanagementunitreferenceMap)
	if err != nil {
		return err
	}
	
	if User, ok := AgentmanagementunitreferenceMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ManagementUnit, ok := AgentmanagementunitreferenceMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := AgentmanagementunitreferenceMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentmanagementunitreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
