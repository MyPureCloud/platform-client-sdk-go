package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nluconfusionmatrixcolumn
type Nluconfusionmatrixcolumn struct { 
	// Name - The name of the intent for the column.
	Name *string `json:"name,omitempty"`


	// Value - The confusion value between the intents
	Value *float32 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
