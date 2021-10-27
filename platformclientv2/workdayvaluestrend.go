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


	// DateReferenceWorkday - The reference workday used to determine the metric definition. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateReferenceWorkday *time.Time `json:"dateReferenceWorkday,omitempty"`


	// Division - The targeted division for the query
	Division *Division `json:"division,omitempty"`


	// User - The targeted user for the query
	User *Userreference `json:"user,omitempty"`


	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`


	// Results - The metric value trends
	Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`


	// PerformanceProfile - The targeted performance profile for the average points
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`


	// Metric - The targeted metric for the average points
	Metric *Addressableentityref `json:"metric,omitempty"`

}

func (o *Workdayvaluestrend) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluestrend
	
	DateStartWorkday := new(string)
	if o.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(o.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if o.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(o.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	
	DateReferenceWorkday := new(string)
	if o.DateReferenceWorkday != nil {
		*DateReferenceWorkday = timeutil.Strftime(o.DateReferenceWorkday, "%Y-%m-%d")
	} else {
		DateReferenceWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		DateReferenceWorkday *string `json:"dateReferenceWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`
		
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		
		Metric *Addressableentityref `json:"metric,omitempty"`
		*Alias
	}{ 
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		DateReferenceWorkday: DateReferenceWorkday,
		
		Division: o.Division,
		
		User: o.User,
		
		Timezone: o.Timezone,
		
		Results: o.Results,
		
		PerformanceProfile: o.PerformanceProfile,
		
		Metric: o.Metric,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdayvaluestrend) UnmarshalJSON(b []byte) error {
	var WorkdayvaluestrendMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdayvaluestrendMap)
	if err != nil {
		return err
	}
	
	if dateStartWorkdayString, ok := WorkdayvaluestrendMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := WorkdayvaluestrendMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if dateReferenceWorkdayString, ok := WorkdayvaluestrendMap["dateReferenceWorkday"].(string); ok {
		DateReferenceWorkday, _ := time.Parse("2006-01-02", dateReferenceWorkdayString)
		o.DateReferenceWorkday = &DateReferenceWorkday
	}
	
	if Division, ok := WorkdayvaluestrendMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if User, ok := WorkdayvaluestrendMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Timezone, ok := WorkdayvaluestrendMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
	
	if Results, ok := WorkdayvaluestrendMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if PerformanceProfile, ok := WorkdayvaluestrendMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Metric, ok := WorkdayvaluestrendMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdayvaluestrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
