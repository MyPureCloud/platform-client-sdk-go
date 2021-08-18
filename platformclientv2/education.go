package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Education) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Education

	
	DateStart := new(string)
	if u.DateStart != nil {
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if u.DateEnd != nil {
		*DateEnd = timeutil.Strftime(u.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	

	return json.Marshal(&struct { 
		School *string `json:"school,omitempty"`
		
		FieldOfStudy *string `json:"fieldOfStudy,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		*Alias
	}{ 
		School: u.School,
		
		FieldOfStudy: u.FieldOfStudy,
		
		Notes: u.Notes,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Education) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
