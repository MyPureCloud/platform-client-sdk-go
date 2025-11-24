package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagenthistoricaladherenceresult
type Wfmagenthistoricaladherenceresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User - The user who submitted the agent historical adherence query
	User *Userreference `json:"user,omitempty"`

	// StartDate - Beginning of the date range that was queried, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
	EndDate *time.Time `json:"endDate,omitempty"`

	// CalculationsCompletedDate - Completed date of calculations that was queried, in ISO-8601 format.
	CalculationsCompletedDate *time.Time `json:"calculationsCompletedDate,omitempty"`

	// TargetAdherencePercentage - Target percentage for this user, in the scale of 0 - 100
	TargetAdherencePercentage *float64 `json:"targetAdherencePercentage,omitempty"`

	// AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`

	// ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`

	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`

	// ExceptionInfo - List of adherence exceptions for this user
	ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`

	// DayMetrics - Adherence and conformance metrics for days in query range
	DayMetrics *[]Agentadherencedaymetrics `json:"dayMetrics,omitempty"`

	// Actuals - List of actual activity with offset for this user
	Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`

	// ScheduledActivities - List of scheduled activities for this user
	ScheduledActivities *[]Agentadherencescheduledactivity `json:"scheduledActivities,omitempty"`

	// SecondaryPresenceLookupItems - List of secondary presence lookup ID to corresponding secondary presence ID item
	SecondaryPresenceLookupItems *[]Secondarypresencelookupitem `json:"secondaryPresenceLookupItems,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagenthistoricaladherenceresult) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagenthistoricaladherenceresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate","CalculationsCompletedDate", }
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
	type Alias Wfmagenthistoricaladherenceresult
	
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
	
	CalculationsCompletedDate := new(string)
	if o.CalculationsCompletedDate != nil {
		
		*CalculationsCompletedDate = timeutil.Strftime(o.CalculationsCompletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CalculationsCompletedDate = nil
	}
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		CalculationsCompletedDate *string `json:"calculationsCompletedDate,omitempty"`
		
		TargetAdherencePercentage *float64 `json:"targetAdherencePercentage,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`
		
		DayMetrics *[]Agentadherencedaymetrics `json:"dayMetrics,omitempty"`
		
		Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`
		
		ScheduledActivities *[]Agentadherencescheduledactivity `json:"scheduledActivities,omitempty"`
		
		SecondaryPresenceLookupItems *[]Secondarypresencelookupitem `json:"secondaryPresenceLookupItems,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		CalculationsCompletedDate: CalculationsCompletedDate,
		
		TargetAdherencePercentage: o.TargetAdherencePercentage,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		
		Impact: o.Impact,
		
		ExceptionInfo: o.ExceptionInfo,
		
		DayMetrics: o.DayMetrics,
		
		Actuals: o.Actuals,
		
		ScheduledActivities: o.ScheduledActivities,
		
		SecondaryPresenceLookupItems: o.SecondaryPresenceLookupItems,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagenthistoricaladherenceresult) UnmarshalJSON(b []byte) error {
	var WfmagenthistoricaladherenceresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagenthistoricaladherenceresultMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmagenthistoricaladherenceresultMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if startDateString, ok := WfmagenthistoricaladherenceresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmagenthistoricaladherenceresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if calculationsCompletedDateString, ok := WfmagenthistoricaladherenceresultMap["calculationsCompletedDate"].(string); ok {
		CalculationsCompletedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", calculationsCompletedDateString)
		o.CalculationsCompletedDate = &CalculationsCompletedDate
	}
	
	if TargetAdherencePercentage, ok := WfmagenthistoricaladherenceresultMap["targetAdherencePercentage"].(float64); ok {
		o.TargetAdherencePercentage = &TargetAdherencePercentage
	}
    
	if AdherencePercentage, ok := WfmagenthistoricaladherenceresultMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := WfmagenthistoricaladherenceresultMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    
	if Impact, ok := WfmagenthistoricaladherenceresultMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if ExceptionInfo, ok := WfmagenthistoricaladherenceresultMap["exceptionInfo"].([]interface{}); ok {
		ExceptionInfoString, _ := json.Marshal(ExceptionInfo)
		json.Unmarshal(ExceptionInfoString, &o.ExceptionInfo)
	}
	
	if DayMetrics, ok := WfmagenthistoricaladherenceresultMap["dayMetrics"].([]interface{}); ok {
		DayMetricsString, _ := json.Marshal(DayMetrics)
		json.Unmarshal(DayMetricsString, &o.DayMetrics)
	}
	
	if Actuals, ok := WfmagenthistoricaladherenceresultMap["actuals"].([]interface{}); ok {
		ActualsString, _ := json.Marshal(Actuals)
		json.Unmarshal(ActualsString, &o.Actuals)
	}
	
	if ScheduledActivities, ok := WfmagenthistoricaladherenceresultMap["scheduledActivities"].([]interface{}); ok {
		ScheduledActivitiesString, _ := json.Marshal(ScheduledActivities)
		json.Unmarshal(ScheduledActivitiesString, &o.ScheduledActivities)
	}
	
	if SecondaryPresenceLookupItems, ok := WfmagenthistoricaladherenceresultMap["secondaryPresenceLookupItems"].([]interface{}); ok {
		SecondaryPresenceLookupItemsString, _ := json.Marshal(SecondaryPresenceLookupItems)
		json.Unmarshal(SecondaryPresenceLookupItemsString, &o.SecondaryPresenceLookupItems)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagenthistoricaladherenceresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
