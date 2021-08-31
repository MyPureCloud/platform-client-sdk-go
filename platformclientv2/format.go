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

func (o *Format) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Format
	
	return json.Marshal(&struct { 
		Flags *[]string `json:"flags,omitempty"`
		*Alias
	}{ 
		Flags: o.Flags,
		Alias:    (*Alias)(o),
	})
}

func (o *Format) UnmarshalJSON(b []byte) error {
	var FormatMap map[string]interface{}
	err := json.Unmarshal(b, &FormatMap)
	if err != nil {
		return err
	}
	
	if Flags, ok := FormatMap["flags"].([]interface{}); ok {
		FlagsString, _ := json.Marshal(Flags)
		json.Unmarshal(FlagsString, &o.Flags)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Format) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
