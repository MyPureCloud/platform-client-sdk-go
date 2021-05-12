package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Tokeninfocloneduser
type Tokeninfocloneduser struct { 
	// Id - User id of the original native user
	Id *string `json:"id,omitempty"`


	// Organization - Organization of the original native user
	Organization *Entity `json:"organization,omitempty"`

}

// String returns a JSON representation of the model
func (o *Tokeninfocloneduser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
