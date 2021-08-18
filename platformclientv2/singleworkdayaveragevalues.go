package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Singleworkdayaveragevalues
type Singleworkdayaveragevalues struct { 
	// DateWorkday - The targeted workday for average value query. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Division - The targeted division for the metrics
	Division *Division `json:"division,omitempty"`


	// User - The targeted user for the metrics
	User *Userreference `json:"user,omitempty"`


	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`


	// Results - The metric value averages
	Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`

}

func (u *Singleworkdayaveragevalues) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Singleworkdayaveragevalues

	
	DateWorkday := new(string)
	if u.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(u.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	

	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Division: u.Division,
		
		User: u.User,
		
		Timezone: u.Timezone,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragevalues) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
