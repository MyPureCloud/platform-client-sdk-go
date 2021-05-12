package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicuserdata
type Adhocrecordingtopicuserdata struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicuserdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
