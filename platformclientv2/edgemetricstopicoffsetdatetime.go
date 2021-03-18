package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicoffsetdatetime
type Edgemetricstopicoffsetdatetime struct { 
	// DateTime
	DateTime *Edgemetricstopiclocaldatetime `json:"dateTime,omitempty"`


	// Offset
	Offset *Edgemetricstopiczoneoffset `json:"offset,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicoffsetdatetime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
