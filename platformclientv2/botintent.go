package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Botintent - A botConnector's bot intention
type Botintent struct { 
	// Name - The name of this intent.  This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
	Name *string `json:"name,omitempty"`


	// Slots - Optional returned data values associated with this intent, limit of 50.
	Slots *map[string]Botslot `json:"slots,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
