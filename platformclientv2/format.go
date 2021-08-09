package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Format
type Format struct { 
	// Flags - The Set of prompt segment format flags i.e. each entry is a part of describing the overall format. E.g. \"format\": { \"flags\": [StringPlayChars] }
	Flags *[]string `json:"flags,omitempty"`

}

// String returns a JSON representation of the model
func (o *Format) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
