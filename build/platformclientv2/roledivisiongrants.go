package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Roledivisiongrants
type Roledivisiongrants struct { 
	// Grants - A list containing pairs of role and division IDs
	Grants *[]Roledivisionpair `json:"grants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Roledivisiongrants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
