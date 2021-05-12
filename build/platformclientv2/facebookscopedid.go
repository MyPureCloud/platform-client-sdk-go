package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookscopedid - Scoped ID for a Facebook user interacting with a page or app
type Facebookscopedid struct { 
	// ScopedId - The unique page/app-specific scopedId for the user
	ScopedId *string `json:"scopedId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookscopedid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
