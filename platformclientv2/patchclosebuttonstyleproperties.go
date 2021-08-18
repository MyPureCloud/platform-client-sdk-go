package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchclosebuttonstyleproperties
type Patchclosebuttonstyleproperties struct { 
	// Color - Color of button. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// Opacity - Opacity of button.
	Opacity *float32 `json:"opacity,omitempty"`

}

func (u *Patchclosebuttonstyleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchclosebuttonstyleproperties

	

	return json.Marshal(&struct { 
		Color *string `json:"color,omitempty"`
		
		Opacity *float32 `json:"opacity,omitempty"`
		*Alias
	}{ 
		Color: u.Color,
		
		Opacity: u.Opacity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchclosebuttonstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
