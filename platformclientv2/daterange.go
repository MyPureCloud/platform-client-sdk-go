package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Daterange
type Daterange struct { 
	// StartDate - The inclusive start of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
	StartDate *string `json:"startDate,omitempty"`


	// EndDate - The inclusive end of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
	EndDate *string `json:"endDate,omitempty"`

}

func (u *Daterange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Daterange

	

	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		*Alias
	}{ 
		StartDate: u.StartDate,
		
		EndDate: u.EndDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Daterange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
