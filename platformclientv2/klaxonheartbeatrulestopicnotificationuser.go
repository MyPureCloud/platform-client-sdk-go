package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatrulestopicnotificationuser
type Klaxonheartbeatrulestopicnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
