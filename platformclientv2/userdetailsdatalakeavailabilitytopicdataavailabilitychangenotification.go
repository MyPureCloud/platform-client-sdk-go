package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { }

func (u *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification

	

	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
