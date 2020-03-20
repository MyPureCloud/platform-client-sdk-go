package platformclientv2
import (
	"encoding/json"
)

// Generateshorttermforecastrequest
type Generateshorttermforecastrequest struct { 
	// Description - Description for the generated forecast.  Pass an empty string for no description
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generateshorttermforecastrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
