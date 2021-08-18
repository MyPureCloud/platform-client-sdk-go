package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingusageresource
type Billingusageresource struct { 
	// Name - Identifies the resource (e.g. license user, device).
	Name *string `json:"name,omitempty"`


	// Date - The date that the usage was first observed by the billing subsystem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`

}

func (u *Billingusageresource) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Billingusageresource

	
	Date := new(string)
	if u.Date != nil {
		
		*Date = timeutil.Strftime(u.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Date *string `json:"date,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Date: Date,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Billingusageresource) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
