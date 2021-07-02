package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scaleasgresponse
type Scaleasgresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DesiredCapacity - Desired capacity for the ASG
	DesiredCapacity *int `json:"desiredCapacity,omitempty"`


	// MinSize - Minimum size for the ASG
	MinSize *int `json:"minSize,omitempty"`


	// MaxSize - Maximum size for the ASG
	MaxSize *int `json:"maxSize,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scaleasgresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
