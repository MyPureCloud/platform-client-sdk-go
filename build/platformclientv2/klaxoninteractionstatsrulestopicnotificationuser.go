package platformclientv2
import (
	"encoding/json"
)

// Klaxoninteractionstatsrulestopicnotificationuser
type Klaxoninteractionstatsrulestopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsrulestopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
