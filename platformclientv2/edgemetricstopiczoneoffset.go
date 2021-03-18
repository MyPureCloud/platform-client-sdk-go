package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopiczoneoffset
type Edgemetricstopiczoneoffset struct { 
	// TotalSeconds
	TotalSeconds *int `json:"totalSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopiczoneoffset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
