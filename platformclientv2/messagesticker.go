package platformclientv2
import (
	"encoding/json"
)

// Messagesticker
type Messagesticker struct { 
	// Url - The location of the sticker, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// Id - The unique id of the the sticker object.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagesticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
