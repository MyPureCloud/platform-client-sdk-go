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

func (o *Userstations) MarshalJSON() ([]byte, error) {
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
		AssociatedStation: o.AssociatedStation,
		
		EffectiveStation: o.EffectiveStation,
		
		DefaultStation: o.DefaultStation,
		
		LastAssociatedStation: o.LastAssociatedStation,
		Alias:    (*Alias)(o),
	})
}

func (o *Userstations) UnmarshalJSON(b []byte) error {
	var UserstationsMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationsMap)
	if err != nil {
		return err
	}
	
	if AssociatedStation, ok := UserstationsMap["associatedStation"].(map[string]interface{}); ok {
		AssociatedStationString, _ := json.Marshal(AssociatedStation)
		json.Unmarshal(AssociatedStationString, &o.AssociatedStation)
	}
	
	if EffectiveStation, ok := UserstationsMap["effectiveStation"].(map[string]interface{}); ok {
		EffectiveStationString, _ := json.Marshal(EffectiveStation)
		json.Unmarshal(EffectiveStationString, &o.EffectiveStation)
	}
	
	if DefaultStation, ok := UserstationsMap["defaultStation"].(map[string]interface{}); ok {
		DefaultStationString, _ := json.Marshal(DefaultStation)
		json.Unmarshal(DefaultStationString, &o.DefaultStation)
	}
	
	if LastAssociatedStation, ok := UserstationsMap["lastAssociatedStation"].(map[string]interface{}); ok {
		LastAssociatedStationString, _ := json.Marshal(LastAssociatedStation)
		json.Unmarshal(LastAssociatedStationString, &o.LastAssociatedStation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
