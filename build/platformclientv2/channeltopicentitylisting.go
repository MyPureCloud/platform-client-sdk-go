package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Channeltopicentitylisting
type Channeltopicentitylisting struct { 
	// Entities
	Entities *[]Channeltopic `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Channeltopicentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
