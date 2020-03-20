package platformclientv2
import (
	"encoding/json"
)

// Listwrapperforecastsourcedaypointer
type Listwrapperforecastsourcedaypointer struct { 
	// Values
	Values *[]Forecastsourcedaypointer `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listwrapperforecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
