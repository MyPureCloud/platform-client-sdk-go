package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluestrend
type Workdayvaluestrend struct { 
	// DateStartWorkday - The start workday for the query range for the metric value trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - The end workday for the query range for the metric value trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// Division - The targeted division for the query
	Division *Division `json:"division,omitempty"`


	// User - The targeted user for the query
	User *Userreference `json:"user,omitempty"`


	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`


	// Results - The metric value trends
	Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`

}

func (u *Workdayvaluestrend) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluestrend

	
	DateStartWorkday := new(string)
	if u.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(u.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if u.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(u.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	

	return json.Marshal(&struct { 
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`
		*Alias
	}{ 
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		Division: u.Division,
		
		User: u.User,
		
		Timezone: u.Timezone,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workdayvaluestrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
