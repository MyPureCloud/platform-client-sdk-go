package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignruleconfigchangeurireference
type Dialercampaignruleconfigchangeurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
