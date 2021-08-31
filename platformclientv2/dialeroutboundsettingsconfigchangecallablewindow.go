package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangecallablewindow
type Dialeroutboundsettingsconfigchangecallablewindow struct { 
	// Mapped
	Mapped *Dialeroutboundsettingsconfigchangeatzmtimeslot `json:"mapped,omitempty"`


	// Unmapped
	Unmapped *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone `json:"unmapped,omitempty"`

}

func (o *Dialeroutboundsettingsconfigchangecallablewindow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangecallablewindow
	
	return json.Marshal(&struct { 
		Mapped *Dialeroutboundsettingsconfigchangeatzmtimeslot `json:"mapped,omitempty"`
		
		Unmapped *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone `json:"unmapped,omitempty"`
		*Alias
	}{ 
		Mapped: o.Mapped,
		
		Unmapped: o.Unmapped,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangecallablewindow) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangecallablewindowMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangecallablewindowMap)
	if err != nil {
		return err
	}
	
	if Mapped, ok := DialeroutboundsettingsconfigchangecallablewindowMap["mapped"].(map[string]interface{}); ok {
		MappedString, _ := json.Marshal(Mapped)
		json.Unmarshal(MappedString, &o.Mapped)
	}
	
	if Unmapped, ok := DialeroutboundsettingsconfigchangecallablewindowMap["unmapped"].(map[string]interface{}); ok {
		UnmappedString, _ := json.Marshal(Unmapped)
		json.Unmarshal(UnmappedString, &o.Unmapped)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangecallablewindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
