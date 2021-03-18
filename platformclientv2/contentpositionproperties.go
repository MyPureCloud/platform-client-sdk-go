package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentpositionproperties
type Contentpositionproperties struct { 
	// Top - Top positioning offset.
	Top *string `json:"top,omitempty"`


	// Bottom - Bottom positioning offset.
	Bottom *string `json:"bottom,omitempty"`


	// Left - Left positioning offset.
	Left *string `json:"left,omitempty"`


	// Right - Right positioning offset.
	Right *string `json:"right,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentpositionproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
