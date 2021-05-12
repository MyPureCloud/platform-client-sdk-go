package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Education
type Education struct { 
	// School
	School *string `json:"school,omitempty"`


	// FieldOfStudy
	FieldOfStudy *string `json:"fieldOfStudy,omitempty"`


	// Notes - Notes about education has a 2000 character limit
	Notes *string `json:"notes,omitempty"`


	// DateStart - Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd - Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

}

// String returns a JSON representation of the model
func (o *Education) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
