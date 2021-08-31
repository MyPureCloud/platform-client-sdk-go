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

func (o *Education) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Education
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%d")
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
		School: o.School,
		
		FieldOfStudy: o.FieldOfStudy,
		
		Notes: o.Notes,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		Alias:    (*Alias)(o),
	})
}

func (o *Education) UnmarshalJSON(b []byte) error {
	var EducationMap map[string]interface{}
	err := json.Unmarshal(b, &EducationMap)
	if err != nil {
		return err
	}
	
	if School, ok := EducationMap["school"].(string); ok {
		o.School = &School
	}
	
	if FieldOfStudy, ok := EducationMap["fieldOfStudy"].(string); ok {
		o.FieldOfStudy = &FieldOfStudy
	}
	
	if Notes, ok := EducationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if dateStartString, ok := EducationMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := EducationMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Education) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
