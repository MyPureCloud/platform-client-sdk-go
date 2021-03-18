package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationstatus
type Gamificationstatus struct { 
	// IsActive - Gamification status of the organization.
	IsActive *bool `json:"isActive,omitempty"`


	// DateStart - Gamification start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

}

// String returns a JSON representation of the model
func (o *Gamificationstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
