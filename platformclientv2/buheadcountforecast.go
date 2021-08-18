package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buheadcountforecast
type Buheadcountforecast struct { 
	// Entities
	Entities *[]Buplanninggroupheadcountforecast `json:"entities,omitempty"`


	// ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

}

func (u *Buheadcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buheadcountforecast

	
	ReferenceStartDate := new(string)
	if u.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(u.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	

	return json.Marshal(&struct { 
		Entities *[]Buplanninggroupheadcountforecast `json:"entities,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		ReferenceStartDate: ReferenceStartDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
