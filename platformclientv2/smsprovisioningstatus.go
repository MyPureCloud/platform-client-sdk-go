package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsprovisioningstatus
type Smsprovisioningstatus struct { 
	// Action - Provisioning action
	Action *string `json:"action,omitempty"`


	// State - Provisioning state
	State *string `json:"state,omitempty"`


	// VarError - Any error associated with a Failed state
	VarError *Errorbody `json:"error,omitempty"`


	// Version - The phone number version associated with the provisioning action
	Version *int `json:"version,omitempty"`

}

func (o *Smsprovisioningstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsprovisioningstatus
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		State: o.State,
		
		VarError: o.VarError,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsprovisioningstatus) UnmarshalJSON(b []byte) error {
	var SmsprovisioningstatusMap map[string]interface{}
	err := json.Unmarshal(b, &SmsprovisioningstatusMap)
	if err != nil {
		return err
	}
	
	if Action, ok := SmsprovisioningstatusMap["action"].(string); ok {
		o.Action = &Action
	}
	
	if State, ok := SmsprovisioningstatusMap["state"].(string); ok {
		o.State = &State
	}
	
	if VarError, ok := SmsprovisioningstatusMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if Version, ok := SmsprovisioningstatusMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsprovisioningstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
