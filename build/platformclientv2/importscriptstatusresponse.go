package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Importscriptstatusresponse
type Importscriptstatusresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Succeeded
	Succeeded *bool `json:"succeeded,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importscriptstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
