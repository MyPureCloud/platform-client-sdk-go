package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignconfigchangeurireference
type Dialercampaignconfigchangeurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
