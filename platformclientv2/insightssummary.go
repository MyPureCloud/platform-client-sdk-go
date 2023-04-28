package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Insightssummary
type Insightssummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entities
	Entities *[]Insightssummaryuseritem `json:"entities,omitempty"`

	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`

	// Total
	Total *int `json:"total,omitempty"`

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

	// PageCount
	PageCount *int `json:"pageCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Insightssummary) SetField(field string, fieldValue interface{}) {
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

func (o Insightssummary) MarshalJSON() ([]byte, error) {
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
	type Alias Insightssummary
	
	return json.Marshal(&struct { 
		Entities *[]Insightssummaryuseritem `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		ComparativePeriod *Workdayperiod `json:"comparativePeriod,omitempty"`
		
		PrimaryPeriod *Workdayperiod `json:"primaryPeriod,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		PerformanceProfile: o.PerformanceProfile,
		
		Division: o.Division,
		
		Granularity: o.Granularity,
		
		ComparativePeriod: o.ComparativePeriod,
		
		PrimaryPeriod: o.PrimaryPeriod,
		
		PageCount: o.PageCount,
		Alias:    (Alias)(o),
	})
}

func (o *Insightssummary) UnmarshalJSON(b []byte) error {
	var InsightssummaryMap map[string]interface{}
	err := json.Unmarshal(b, &InsightssummaryMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := InsightssummaryMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := InsightssummaryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := InsightssummaryMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := InsightssummaryMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PerformanceProfile, ok := InsightssummaryMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Division, ok := InsightssummaryMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Granularity, ok := InsightssummaryMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if ComparativePeriod, ok := InsightssummaryMap["comparativePeriod"].(map[string]interface{}); ok {
		ComparativePeriodString, _ := json.Marshal(ComparativePeriod)
		json.Unmarshal(ComparativePeriodString, &o.ComparativePeriod)
	}
	
	if PrimaryPeriod, ok := InsightssummaryMap["primaryPeriod"].(map[string]interface{}); ok {
		PrimaryPeriodString, _ := json.Marshal(PrimaryPeriod)
		json.Unmarshal(PrimaryPeriodString, &o.PrimaryPeriod)
	}
	
	if PageCount, ok := InsightssummaryMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Insightssummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
