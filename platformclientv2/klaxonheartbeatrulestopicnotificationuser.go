package platformclientv2
import (
	"encoding/json"
)

// Klaxonheartbeatrulestopicnotificationuser
type Klaxonheartbeatrulestopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
