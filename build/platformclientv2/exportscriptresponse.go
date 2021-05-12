package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Exportscriptresponse
type Exportscriptresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`

}

// String returns a JSON representation of the model
func (o *Exportscriptresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
