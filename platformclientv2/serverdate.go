package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Serverdate
type Serverdate struct { 
	// CurrentDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CurrentDate *time.Time `json:"currentDate,omitempty"`

}

func (u *Serverdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Serverdate

	
	CurrentDate := new(string)
	if u.CurrentDate != nil {
		
		*CurrentDate = timeutil.Strftime(u.CurrentDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CurrentDate = nil
	}
	

	return json.Marshal(&struct { 
		CurrentDate *string `json:"currentDate,omitempty"`
		*Alias
	}{ 
		CurrentDate: CurrentDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Serverdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
