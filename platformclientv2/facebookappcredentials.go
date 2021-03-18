package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookappcredentials
type Facebookappcredentials struct { 
	// Id - Genesys Cloud Facebook App Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookappcredentials) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
