package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult
type Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User
	User *Wfmhistoricaladherenceagentcalculationscompletetopicuserreference `json:"user,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

	// CalculationsCompletedDate
	CalculationsCompletedDate *time.Time `json:"calculationsCompletedDate,omitempty"`

	// TargetAdherencePercentage
	TargetAdherencePercentage *float32 `json:"targetAdherencePercentage,omitempty"`

	// AdherencePercentage
	AdherencePercentage *float32 `json:"adherencePercentage,omitempty"`

	// ConformancePercentage
	ConformancePercentage *float32 `json:"conformancePercentage,omitempty"`

	// Impact
	Impact *string `json:"impact,omitempty"`

	// ExceptionInfo
	ExceptionInfo *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`

	// DayMetrics
	DayMetrics *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentdaymetrics `json:"dayMetrics,omitempty"`

	// Actuals
	Actuals *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceactuals `json:"actuals,omitempty"`

	// ScheduledActivities
	ScheduledActivities *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherencescheduledactivity `json:"scheduledActivities,omitempty"`

	// SecondaryPresenceLookupItems
	SecondaryPresenceLookupItems *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmsecondarypresencelookupitem `json:"secondaryPresenceLookupItems,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult) SetField(field string, fieldValue interface{}) {
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

func (o Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult
	
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
		User *Wfmhistoricaladherenceagentcalculationscompletetopicuserreference `json:"user,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		CalculationsCompletedDate *string `json:"calculationsCompletedDate,omitempty"`
		
		TargetAdherencePercentage *float32 `json:"targetAdherencePercentage,omitempty"`
		
		AdherencePercentage *float32 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float32 `json:"conformancePercentage,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		ExceptionInfo *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`
		
		DayMetrics *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentdaymetrics `json:"dayMetrics,omitempty"`
		
		Actuals *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceactuals `json:"actuals,omitempty"`
		
		ScheduledActivities *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherencescheduledactivity `json:"scheduledActivities,omitempty"`
		
		SecondaryPresenceLookupItems *[]Wfmhistoricaladherenceagentcalculationscompletetopicwfmsecondarypresencelookupitem `json:"secondaryPresenceLookupItems,omitempty"`
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

func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if startDateString, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if calculationsCompletedDateString, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["calculationsCompletedDate"].(string); ok {
		CalculationsCompletedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", calculationsCompletedDateString)
		o.CalculationsCompletedDate = &CalculationsCompletedDate
	}
	
	if TargetAdherencePercentage, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["targetAdherencePercentage"].(float64); ok {
		TargetAdherencePercentageFloat32 := float32(TargetAdherencePercentage)
		o.TargetAdherencePercentage = &TargetAdherencePercentageFloat32
	}
    
	if AdherencePercentage, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["adherencePercentage"].(float64); ok {
		AdherencePercentageFloat32 := float32(AdherencePercentage)
		o.AdherencePercentage = &AdherencePercentageFloat32
	}
    
	if ConformancePercentage, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["conformancePercentage"].(float64); ok {
		ConformancePercentageFloat32 := float32(ConformancePercentage)
		o.ConformancePercentage = &ConformancePercentageFloat32
	}
    
	if Impact, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if ExceptionInfo, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["exceptionInfo"].([]interface{}); ok {
		ExceptionInfoString, _ := json.Marshal(ExceptionInfo)
		json.Unmarshal(ExceptionInfoString, &o.ExceptionInfo)
	}
	
	if DayMetrics, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["dayMetrics"].([]interface{}); ok {
		DayMetricsString, _ := json.Marshal(DayMetrics)
		json.Unmarshal(DayMetricsString, &o.DayMetrics)
	}
	
	if Actuals, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["actuals"].([]interface{}); ok {
		ActualsString, _ := json.Marshal(Actuals)
		json.Unmarshal(ActualsString, &o.Actuals)
	}
	
	if ScheduledActivities, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["scheduledActivities"].([]interface{}); ok {
		ScheduledActivitiesString, _ := json.Marshal(ScheduledActivities)
		json.Unmarshal(ScheduledActivitiesString, &o.ScheduledActivities)
	}
	
	if SecondaryPresenceLookupItems, ok := WfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresultMap["secondaryPresenceLookupItems"].([]interface{}); ok {
		SecondaryPresenceLookupItemsString, _ := json.Marshal(SecondaryPresenceLookupItems)
		json.Unmarshal(SecondaryPresenceLookupItemsString, &o.SecondaryPresenceLookupItems)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceagentcalculationscompletetopicwfmhistoricaladherenceagentresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
