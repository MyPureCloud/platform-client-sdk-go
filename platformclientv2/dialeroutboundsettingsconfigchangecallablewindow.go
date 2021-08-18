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

func (u *Dialeroutboundsettingsconfigchangecallablewindow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangecallablewindow

	

	return json.Marshal(&struct { 
		Mapped *Dialeroutboundsettingsconfigchangeatzmtimeslot `json:"mapped,omitempty"`
		
		Unmapped *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone `json:"unmapped,omitempty"`
		*Alias
	}{ 
		Mapped: u.Mapped,
		
		Unmapped: u.Unmapped,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangecallablewindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
