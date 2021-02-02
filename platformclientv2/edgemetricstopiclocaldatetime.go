package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopiclocaldatetime
type Edgemetricstopiclocaldatetime struct { 
	// Date
	Date *Edgemetricstopiclocaldate `json:"date,omitempty"`


	// Time
	Time *Edgemetricstopiclocaltime `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopiclocaldatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
