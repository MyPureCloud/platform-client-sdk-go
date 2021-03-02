package platformclientv2
import (
	"encoding/json"
)

// Trunkinstancetopictrunkmetricsnetworktypeip
type Trunkinstancetopictrunkmetricsnetworktypeip struct { 
	// Address
	Address *string `json:"address,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
