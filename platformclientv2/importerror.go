package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Importerror
type Importerror struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Line
	Line *int `json:"line,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
