package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Outofofficeeventoutofoffice
type Outofofficeeventoutofoffice struct { 
	// User
	User *Outofofficeeventuser `json:"user,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// Indefinite
	Indefinite *bool `json:"indefinite,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outofofficeeventoutofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
