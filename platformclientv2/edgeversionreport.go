package platformclientv2
import (
	"encoding/json"
)

// Edgeversionreport
type Edgeversionreport struct { 
	// OldestVersion
	OldestVersion *Edgeversioninformation `json:"oldestVersion,omitempty"`


	// NewestVersion
	NewestVersion *Edgeversioninformation `json:"newestVersion,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeversionreport) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
