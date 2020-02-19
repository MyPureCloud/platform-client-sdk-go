package platformclientv2
import (
	"encoding/json"
)

// Orgwhitelistsettings
type Orgwhitelistsettings struct { 
	// EnableWhitelist
	EnableWhitelist *bool `json:"enableWhitelist,omitempty"`


	// DomainWhitelist
	DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`

}

// String returns a JSON representation of the model
func (o *Orgwhitelistsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
