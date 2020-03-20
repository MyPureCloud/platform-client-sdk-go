package platformclientv2
import (
	"encoding/json"
)

// Exportscriptresponse
type Exportscriptresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`

}

// String returns a JSON representation of the model
func (o *Exportscriptresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
