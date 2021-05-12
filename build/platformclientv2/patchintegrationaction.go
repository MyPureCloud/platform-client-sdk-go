package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchintegrationaction
type Patchintegrationaction struct { 
	// Id - ID of the integration action to be invoked.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchintegrationaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
