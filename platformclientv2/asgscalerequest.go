package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Asgscalerequest
type Asgscalerequest struct { 
	// DesiredCapacity - The desired capacity of the ASG
	DesiredCapacity *int `json:"desiredCapacity,omitempty"`


	// MinimumCapacity - The minimum capacity of the ASG
	MinimumCapacity *int `json:"minimumCapacity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Asgscalerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
