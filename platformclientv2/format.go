package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Format
type Format struct { 
	// Flags - The Set of prompt segment format flags i.e. each entry is a part of describing the overall format. E.g. \"format\": { \"flags\": [StringPlayChars] }
	Flags *[]string `json:"flags,omitempty"`

}

func (u *Format) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Format

	

	return json.Marshal(&struct { 
		Flags *[]string `json:"flags,omitempty"`
		*Alias
	}{ 
		Flags: u.Flags,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Format) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
