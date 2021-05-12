package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangecallablewindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
