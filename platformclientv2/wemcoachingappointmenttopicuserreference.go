package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopicuserreference
type Wemcoachingappointmenttopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopicuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
