package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicoutofoffice
type Agentactivitychangedtopicoutofoffice struct { 
	// Active
	Active *bool `json:"active,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicoutofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
