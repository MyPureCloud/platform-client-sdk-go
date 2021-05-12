package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapschedulegroup
type Actionmapschedulegroup struct { 
	// Id - The ID of the action maps's associated schedule group.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
