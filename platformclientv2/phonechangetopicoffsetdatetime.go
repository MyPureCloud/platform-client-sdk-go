package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicoffsetdatetime
type Phonechangetopicoffsetdatetime struct { 
	// DateTime
	DateTime *Phonechangetopiclocaldatetime `json:"dateTime,omitempty"`


	// Offset
	Offset *Phonechangetopiczoneoffset `json:"offset,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopicoffsetdatetime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
