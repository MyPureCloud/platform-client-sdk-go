package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopiclocaldatetime
type Phonechangetopiclocaldatetime struct { 
	// Date
	Date *Phonechangetopiclocaldate `json:"date,omitempty"`


	// Time
	Time *Phonechangetopiclocaltime `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopiclocaldatetime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
