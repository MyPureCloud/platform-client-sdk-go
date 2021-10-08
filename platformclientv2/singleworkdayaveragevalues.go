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


	// PerformanceProfile - The targeted performance profile for the average points
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`

}

func (o *Singleworkdayaveragevalues) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Singleworkdayaveragevalues
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`
		
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Division: o.Division,
		
		User: o.User,
		
		Timezone: o.Timezone,
		
		Results: o.Results,
		
		PerformanceProfile: o.PerformanceProfile,
		Alias:    (*Alias)(o),
	})
}

func (o *Singleworkdayaveragevalues) UnmarshalJSON(b []byte) error {
	var SingleworkdayaveragevaluesMap map[string]interface{}
	err := json.Unmarshal(b, &SingleworkdayaveragevaluesMap)
	if err != nil {
		return err
	}
	
	if dateWorkdayString, ok := SingleworkdayaveragevaluesMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	
	if Division, ok := SingleworkdayaveragevaluesMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if User, ok := SingleworkdayaveragevaluesMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Timezone, ok := SingleworkdayaveragevaluesMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
	
	if Results, ok := SingleworkdayaveragevaluesMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if PerformanceProfile, ok := SingleworkdayaveragevaluesMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragevalues) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
