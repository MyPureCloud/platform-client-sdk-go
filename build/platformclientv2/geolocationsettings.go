package platformclientv2
import (
	"encoding/json"
)

// Geolocationsettings
type Geolocationsettings struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MapboxKey
	MapboxKey *string `json:"mapboxKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Geolocationsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
