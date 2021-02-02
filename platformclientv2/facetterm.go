package platformclientv2
import (
	"time"
	"encoding/json"
)

// Facetterm
type Facetterm struct { 
	// Term
	Term *string `json:"term,omitempty"`


	// Key
	Key *int `json:"key,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`


	// Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facetterm) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
