package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Insightsdetailsmetricperiodpoints
type Insightsdetailsmetricperiodpoints struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Points - Points scored
	Points *int `json:"points,omitempty"`

	// MaxPoints - Max possible points
	MaxPoints *int `json:"maxPoints,omitempty"`

	// DataPointCount - Number of data points
	DataPointCount *int `json:"dataPointCount,omitempty"`

	// PercentOfGoal - Percentage of the goal
	PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`

	// RankTotalPoints - The agent's rank in leader board for points on this metric
	RankTotalPoints *int `json:"rankTotalPoints,omitempty"`

	// RankPercentagePoints - The agent's rank in leader board for percentage on this metric
	RankPercentagePoints *int `json:"rankPercentagePoints,omitempty"`

	// Value - Average value
	Value *float64 `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Insightsdetailsmetricperiodpoints) SetField(field string, fieldValue interface{}) {
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

func (o Insightsdetailsmetricperiodpoints) MarshalJSON() ([]byte, error) {
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
	type Alias Insightsdetailsmetricperiodpoints
	
	return json.Marshal(&struct { 
		Points *int `json:"points,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		DataPointCount *int `json:"dataPointCount,omitempty"`
		
		PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`
		
		RankTotalPoints *int `json:"rankTotalPoints,omitempty"`
		
		RankPercentagePoints *int `json:"rankPercentagePoints,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		Alias
	}{ 
		Points: o.Points,
		
		MaxPoints: o.MaxPoints,
		
		DataPointCount: o.DataPointCount,
		
		PercentOfGoal: o.PercentOfGoal,
		
		RankTotalPoints: o.RankTotalPoints,
		
		RankPercentagePoints: o.RankPercentagePoints,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Insightsdetailsmetricperiodpoints) UnmarshalJSON(b []byte) error {
	var InsightsdetailsmetricperiodpointsMap map[string]interface{}
	err := json.Unmarshal(b, &InsightsdetailsmetricperiodpointsMap)
	if err != nil {
		return err
	}
	
	if Points, ok := InsightsdetailsmetricperiodpointsMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if MaxPoints, ok := InsightsdetailsmetricperiodpointsMap["maxPoints"].(float64); ok {
		MaxPointsInt := int(MaxPoints)
		o.MaxPoints = &MaxPointsInt
	}
	
	if DataPointCount, ok := InsightsdetailsmetricperiodpointsMap["dataPointCount"].(float64); ok {
		DataPointCountInt := int(DataPointCount)
		o.DataPointCount = &DataPointCountInt
	}
	
	if PercentOfGoal, ok := InsightsdetailsmetricperiodpointsMap["percentOfGoal"].(float64); ok {
		o.PercentOfGoal = &PercentOfGoal
	}
    
	if RankTotalPoints, ok := InsightsdetailsmetricperiodpointsMap["rankTotalPoints"].(float64); ok {
		RankTotalPointsInt := int(RankTotalPoints)
		o.RankTotalPoints = &RankTotalPointsInt
	}
	
	if RankPercentagePoints, ok := InsightsdetailsmetricperiodpointsMap["rankPercentagePoints"].(float64); ok {
		RankPercentagePointsInt := int(RankPercentagePoints)
		o.RankPercentagePoints = &RankPercentagePointsInt
	}
	
	if Value, ok := InsightsdetailsmetricperiodpointsMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Insightsdetailsmetricperiodpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
