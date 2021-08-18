package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Singleworkdayaveragepoints
type Singleworkdayaveragepoints struct { 
	// DateWorkday - Queried target workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Division - The targeted division for the average points
	Division *Division `json:"division,omitempty"`


	// AveragePoints - The average points per agent earned within the division
	AveragePoints *float64 `json:"averagePoints,omitempty"`

}

func (u *Singleworkdayaveragepoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Singleworkdayaveragepoints

	
	DateWorkday := new(string)
	if u.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(u.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	

	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		AveragePoints *float64 `json:"averagePoints,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Division: u.Division,
		
		AveragePoints: u.AveragePoints,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragepoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
