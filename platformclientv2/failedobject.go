package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Failedobject
type Failedobject struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

}

// String returns a JSON representation of the model
func (o *Failedobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
