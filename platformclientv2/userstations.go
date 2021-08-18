package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Userstations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstations

	

	return json.Marshal(&struct { 
		AssociatedStation *Userstation `json:"associatedStation,omitempty"`
		
		EffectiveStation *Userstation `json:"effectiveStation,omitempty"`
		
		DefaultStation *Userstation `json:"defaultStation,omitempty"`
		
		LastAssociatedStation *Userstation `json:"lastAssociatedStation,omitempty"`
		*Alias
	}{ 
		AssociatedStation: u.AssociatedStation,
		
		EffectiveStation: u.EffectiveStation,
		
		DefaultStation: u.DefaultStation,
		
		LastAssociatedStation: u.LastAssociatedStation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
