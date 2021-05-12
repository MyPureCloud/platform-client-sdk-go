package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataavailabilityresponse
type Dataavailabilityresponse struct { 
	// DataAvailabilityDate - Date and time before which data is guaranteed to be available in the datalake. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dataavailabilityresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
