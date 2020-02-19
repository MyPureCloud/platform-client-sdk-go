package platformclientv2
import (
	"encoding/json"
)

// Userstations
type Userstations struct { 
	// AssociatedStation - Current associated station for this user.
	AssociatedStation *Userstation `json:"associatedStation,omitempty"`


	// EffectiveStation - The station where the user can be reached based on their default and associated station.
	EffectiveStation *Userstation `json:"effectiveStation,omitempty"`


	// DefaultStation - Default station to be used if not associated with a station.
	DefaultStation *Userstation `json:"defaultStation,omitempty"`


	// LastAssociatedStation - Last associated station for this user.
	LastAssociatedStation *Userstation `json:"lastAssociatedStation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstations) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
