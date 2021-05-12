package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Setuuidatarequest
type Setuuidatarequest struct { 
	// UuiData - The value of the uuiData to set.
	UuiData *string `json:"uuiData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Setuuidatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
