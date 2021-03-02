package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopiclocaldatetime
type Trunkmetricstopiclocaldatetime struct { 
	// Date
	Date *Trunkmetricstopiclocaldate `json:"date,omitempty"`


	// Time
	Time *Trunkmetricstopiclocaltime `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopiclocaldatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
