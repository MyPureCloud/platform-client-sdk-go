package platformclientv2
import (
	"encoding/json"
)

// Userscheduleadherencelisting
type Userscheduleadherencelisting struct { 
	// Entities
	Entities *[]Userscheduleadherence `json:"entities,omitempty"`


	// DownloadUrl - The downloadUrl if the response is too large to send directly via http response
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userscheduleadherencelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
