package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcontentofferstyleproperties
type Patchcontentofferstyleproperties struct { 
	// Padding - Padding of the offer. (eg. 10px)
	Padding *string `json:"padding,omitempty"`


	// Color - Text color of the offer. (eg. #FF0000)
	Color *string `json:"color,omitempty"`


	// BackgroundColor - Background color of the offer. (eg. #000000)
	BackgroundColor *string `json:"backgroundColor,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchcontentofferstyleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
