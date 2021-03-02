package platformclientv2
import (
	"encoding/json"
)

// Phonechangetopiclocaldatetime
type Phonechangetopiclocaldatetime struct { 
	// Date
	Date *Phonechangetopiclocaldate `json:"date,omitempty"`


	// Time
	Time *Phonechangetopiclocaltime `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopiclocaldatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
