package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Category - List of available Action categories.
type Category struct { 
	// Name - Category name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Category) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
