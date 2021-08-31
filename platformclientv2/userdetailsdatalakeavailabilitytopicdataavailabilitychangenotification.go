package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { }

func (o *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) UnmarshalJSON(b []byte) error {
	var UserdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &UserdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
