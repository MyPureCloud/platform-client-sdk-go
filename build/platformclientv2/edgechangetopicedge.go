package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgechangetopicedge
type Edgechangetopicedge struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// OnlineStatus
	OnlineStatus *string `json:"onlineStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgechangetopicedge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
