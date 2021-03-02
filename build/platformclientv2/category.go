package platformclientv2
import (
	"encoding/json"
)

// Category - List of available Action categories.
type Category struct { 
	// Name - Category name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Category) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
