package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicworkspacedata
type Adhocrecordingtopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
