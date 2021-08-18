package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthlasttokenissued
type Oauthlasttokenissued struct { 
	// DateIssued - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateIssued *time.Time `json:"dateIssued,omitempty"`

}

func (u *Oauthlasttokenissued) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthlasttokenissued

	
	DateIssued := new(string)
	if u.DateIssued != nil {
		
		*DateIssued = timeutil.Strftime(u.DateIssued, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIssued = nil
	}
	

	return json.Marshal(&struct { 
		DateIssued *string `json:"dateIssued,omitempty"`
		*Alias
	}{ 
		DateIssued: DateIssued,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Oauthlasttokenissued) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
