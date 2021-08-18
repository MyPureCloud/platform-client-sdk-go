package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataavailabilityresponse
type Dataavailabilityresponse struct { 
	// DataAvailabilityDate - Date and time before which data is guaranteed to be available in the datalake. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`

}

func (u *Dataavailabilityresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataavailabilityresponse

	
	DataAvailabilityDate := new(string)
	if u.DataAvailabilityDate != nil {
		
		*DataAvailabilityDate = timeutil.Strftime(u.DataAvailabilityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DataAvailabilityDate = nil
	}
	

	return json.Marshal(&struct { 
		DataAvailabilityDate *string `json:"dataAvailabilityDate,omitempty"`
		*Alias
	}{ 
		DataAvailabilityDate: DataAvailabilityDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dataavailabilityresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
