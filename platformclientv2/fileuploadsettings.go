package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Fileuploadsettings - File upload settings for messenger
type Fileuploadsettings struct { 
	// Modes - The list of supported file upload modes
	Modes *[]Fileuploadmode `json:"modes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Fileuploadsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
