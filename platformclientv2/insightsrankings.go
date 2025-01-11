package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Insightsrankings
type Insightsrankings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PerformanceProfile - The performance profile
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`

	// Division - The division
	Division *Divisionreference `json:"division,omitempty"`

	// Granularity - Granularity
	Granularity *string `json:"granularity,omitempty"`

	// ComparativePeriod - The comparative period work day date range
	ComparativePeriod *Workdayperiod `json:"comparativePeriod,omitempty"`

	// PrimaryPeriod - The primary period work day date range
	PrimaryPeriod *Workdayperiod `json:"primaryPeriod,omitempty"`

	// Leaders - The leaders users
	Leaders *[]Insightssummaryuseritem `json:"leaders,omitempty"`

	// Trailers - The trailing users
	Trailers *[]Insightssummaryuseritem `json:"trailers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Insightsrankings) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Insightsrankings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Insightsrankings
	
	return json.Marshal(&struct { 
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		ComparativePeriod *Workdayperiod `json:"comparativePeriod,omitempty"`
		
		PrimaryPeriod *Workdayperiod `json:"primaryPeriod,omitempty"`
		
		Leaders *[]Insightssummaryuseritem `json:"leaders,omitempty"`
		
		Trailers *[]Insightssummaryuseritem `json:"trailers,omitempty"`
		Alias
	}{ 
		PerformanceProfile: o.PerformanceProfile,
		
		Division: o.Division,
		
		Granularity: o.Granularity,
		
		ComparativePeriod: o.ComparativePeriod,
		
		PrimaryPeriod: o.PrimaryPeriod,
		
		Leaders: o.Leaders,
		
		Trailers: o.Trailers,
		Alias:    (Alias)(o),
	})
}

func (o *Insightsrankings) UnmarshalJSON(b []byte) error {
	var InsightsrankingsMap map[string]interface{}
	err := json.Unmarshal(b, &InsightsrankingsMap)
	if err != nil {
		return err
	}
	
	if PerformanceProfile, ok := InsightsrankingsMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Division, ok := InsightsrankingsMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Granularity, ok := InsightsrankingsMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if ComparativePeriod, ok := InsightsrankingsMap["comparativePeriod"].(map[string]interface{}); ok {
		ComparativePeriodString, _ := json.Marshal(ComparativePeriod)
		json.Unmarshal(ComparativePeriodString, &o.ComparativePeriod)
	}
	
	if PrimaryPeriod, ok := InsightsrankingsMap["primaryPeriod"].(map[string]interface{}); ok {
		PrimaryPeriodString, _ := json.Marshal(PrimaryPeriod)
		json.Unmarshal(PrimaryPeriodString, &o.PrimaryPeriod)
	}
	
	if Leaders, ok := InsightsrankingsMap["leaders"].([]interface{}); ok {
		LeadersString, _ := json.Marshal(Leaders)
		json.Unmarshal(LeadersString, &o.Leaders)
	}
	
	if Trailers, ok := InsightsrankingsMap["trailers"].([]interface{}); ok {
		TrailersString, _ := json.Marshal(Trailers)
		json.Unmarshal(TrailersString, &o.Trailers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Insightsrankings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
