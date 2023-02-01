package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradayresponse
type Buintradayresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`

	// IntervalLengthMinutes - The aggregation period in minutes, which determines the interval duration of the returned data
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

	// NoDataReason - If not null, the reason there was no data for the request
	NoDataReason *string `json:"noDataReason,omitempty"`

	// Categories - The categories to which this data corresponds
	Categories *[]string `json:"categories,omitempty"`

	// ShortTermForecast - Short term forecast reference
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`

	// Schedule - Schedule reference
	Schedule *Buschedulereference `json:"schedule,omitempty"`

	// IntradayDataGroupings - Intraday data grouped by a single media type and set of planning group IDs
	IntradayDataGroupings *[]Buintradaydatagroup `json:"intradayDataGroupings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buintradayresponse) SetField(field string, fieldValue interface{}) {
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

func (o Buintradayresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
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
	type Alias Buintradayresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		
		IntradayDataGroupings *[]Buintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		NoDataReason: o.NoDataReason,
		
		Categories: o.Categories,
		
		ShortTermForecast: o.ShortTermForecast,
		
		Schedule: o.Schedule,
		
		IntradayDataGroupings: o.IntradayDataGroupings,
		Alias:    (Alias)(o),
	})
}

func (o *Buintradayresponse) UnmarshalJSON(b []byte) error {
	var BuintradayresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuintradayresponseMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuintradayresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BuintradayresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if IntervalLengthMinutes, ok := BuintradayresponseMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if NoDataReason, ok := BuintradayresponseMap["noDataReason"].(string); ok {
		o.NoDataReason = &NoDataReason
	}
    
	if Categories, ok := BuintradayresponseMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if ShortTermForecast, ok := BuintradayresponseMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if Schedule, ok := BuintradayresponseMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if IntradayDataGroupings, ok := BuintradayresponseMap["intradayDataGroupings"].([]interface{}); ok {
		IntradayDataGroupingsString, _ := json.Marshal(IntradayDataGroupings)
		json.Unmarshal(IntradayDataGroupingsString, &o.IntradayDataGroupings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buintradayresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
