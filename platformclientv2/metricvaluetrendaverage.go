package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Metricvaluetrendaverage
type Metricvaluetrendaverage struct { 
	// DateStartWorkday - The targeted start workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - The targeted end workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// DateReferenceWorkday - The targeted reference workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateReferenceWorkday *time.Time `json:"dateReferenceWorkday,omitempty"`


	// Division - The targeted division for the metrics
	Division *Division `json:"division,omitempty"`


	// User - The targeted user for the metrics
	User *Userreference `json:"user,omitempty"`


	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`


	// Result - The metric value trend and average
	Result *Workdayvaluesmetricitem `json:"result,omitempty"`


	// PerformanceProfile - The targeted performance profile for the average points
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`


	// Metric - The targeted performance profile for the average points
	Metric *Addressableentityref `json:"metric,omitempty"`

}

func (o *Metricvaluetrendaverage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metricvaluetrendaverage
	
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
		
		Result *Workdayvaluesmetricitem `json:"result,omitempty"`
		
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
		
		Result: o.Result,
		
		PerformanceProfile: o.PerformanceProfile,
		
		Metric: o.Metric,
		Alias:    (*Alias)(o),
	})
}

func (o *Metricvaluetrendaverage) UnmarshalJSON(b []byte) error {
	var MetricvaluetrendaverageMap map[string]interface{}
	err := json.Unmarshal(b, &MetricvaluetrendaverageMap)
	if err != nil {
		return err
	}
	
	if dateStartWorkdayString, ok := MetricvaluetrendaverageMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := MetricvaluetrendaverageMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if dateReferenceWorkdayString, ok := MetricvaluetrendaverageMap["dateReferenceWorkday"].(string); ok {
		DateReferenceWorkday, _ := time.Parse("2006-01-02", dateReferenceWorkdayString)
		o.DateReferenceWorkday = &DateReferenceWorkday
	}
	
	if Division, ok := MetricvaluetrendaverageMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if User, ok := MetricvaluetrendaverageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Timezone, ok := MetricvaluetrendaverageMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
    
	if Result, ok := MetricvaluetrendaverageMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if PerformanceProfile, ok := MetricvaluetrendaverageMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Metric, ok := MetricvaluetrendaverageMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Metricvaluetrendaverage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
