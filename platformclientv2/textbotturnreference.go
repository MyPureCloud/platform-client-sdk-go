package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotturnreference - A reference to a bot flow turn.
type Textbotturnreference struct { 
	// Id - The id of the turn.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotturnreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
