package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageresult
type Historicalshrinkageresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - Beginning of the date range that was queried, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
	EndDate *time.Time `json:"endDate,omitempty"`

	// TotalScheduledDurationSeconds - Total duration in seconds for which agents in the management unit are scheduled
	TotalScheduledDurationSeconds *int `json:"totalScheduledDurationSeconds,omitempty"`

	// TotalLoggedInDurationSeconds - Total duration in seconds for which agents in the management unit are actually logged-in
	TotalLoggedInDurationSeconds *int `json:"totalLoggedInDurationSeconds,omitempty"`

	// AggregatedShrinkage - Aggregated shrinkage data for all the activity categories
	AggregatedShrinkage *Historicalshrinkageaggregateresponse `json:"aggregatedShrinkage,omitempty"`

	// ShrinkageForActivityCategories - Shrinkage for activity categories
	ShrinkageForActivityCategories *[]Historicalshrinkageactivitycategoryresponse `json:"shrinkageForActivityCategories,omitempty"`

	// BusinessUnitIds - List of all business units of all the agents in response
	BusinessUnitIds *[]string `json:"businessUnitIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicalshrinkageresult) SetField(field string, fieldValue interface{}) {
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

func (o Historicalshrinkageresult) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Historicalshrinkageresult
	
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
		
		TotalScheduledDurationSeconds *int `json:"totalScheduledDurationSeconds,omitempty"`
		
		TotalLoggedInDurationSeconds *int `json:"totalLoggedInDurationSeconds,omitempty"`
		
		AggregatedShrinkage *Historicalshrinkageaggregateresponse `json:"aggregatedShrinkage,omitempty"`
		
		ShrinkageForActivityCategories *[]Historicalshrinkageactivitycategoryresponse `json:"shrinkageForActivityCategories,omitempty"`
		
		BusinessUnitIds *[]string `json:"businessUnitIds,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		TotalScheduledDurationSeconds: o.TotalScheduledDurationSeconds,
		
		TotalLoggedInDurationSeconds: o.TotalLoggedInDurationSeconds,
		
		AggregatedShrinkage: o.AggregatedShrinkage,
		
		ShrinkageForActivityCategories: o.ShrinkageForActivityCategories,
		
		BusinessUnitIds: o.BusinessUnitIds,
		Alias:    (Alias)(o),
	})
}

func (o *Historicalshrinkageresult) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageresultMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageresultMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := HistoricalshrinkageresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := HistoricalshrinkageresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if TotalScheduledDurationSeconds, ok := HistoricalshrinkageresultMap["totalScheduledDurationSeconds"].(float64); ok {
		TotalScheduledDurationSecondsInt := int(TotalScheduledDurationSeconds)
		o.TotalScheduledDurationSeconds = &TotalScheduledDurationSecondsInt
	}
	
	if TotalLoggedInDurationSeconds, ok := HistoricalshrinkageresultMap["totalLoggedInDurationSeconds"].(float64); ok {
		TotalLoggedInDurationSecondsInt := int(TotalLoggedInDurationSeconds)
		o.TotalLoggedInDurationSeconds = &TotalLoggedInDurationSecondsInt
	}
	
	if AggregatedShrinkage, ok := HistoricalshrinkageresultMap["aggregatedShrinkage"].(map[string]interface{}); ok {
		AggregatedShrinkageString, _ := json.Marshal(AggregatedShrinkage)
		json.Unmarshal(AggregatedShrinkageString, &o.AggregatedShrinkage)
	}
	
	if ShrinkageForActivityCategories, ok := HistoricalshrinkageresultMap["shrinkageForActivityCategories"].([]interface{}); ok {
		ShrinkageForActivityCategoriesString, _ := json.Marshal(ShrinkageForActivityCategories)
		json.Unmarshal(ShrinkageForActivityCategoriesString, &o.ShrinkageForActivityCategories)
	}
	
	if BusinessUnitIds, ok := HistoricalshrinkageresultMap["businessUnitIds"].([]interface{}); ok {
		BusinessUnitIdsString, _ := json.Marshal(BusinessUnitIds)
		json.Unmarshal(BusinessUnitIdsString, &o.BusinessUnitIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
