package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Groupgreetingeventgreetingowner
type Groupgreetingeventgreetingowner struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingowner) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
