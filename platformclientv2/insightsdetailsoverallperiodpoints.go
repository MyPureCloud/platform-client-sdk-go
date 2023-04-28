package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Insightsdetailsoverallperiodpoints
type Insightsdetailsoverallperiodpoints struct { 
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
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Insightsdetailsoverallperiodpoints) SetField(field string, fieldValue interface{}) {
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

func (o Insightsdetailsoverallperiodpoints) MarshalJSON() ([]byte, error) {
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
	type Alias Insightsdetailsoverallperiodpoints
	
	return json.Marshal(&struct { 
		Points *int `json:"points,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		DataPointCount *int `json:"dataPointCount,omitempty"`
		
		PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`
		
		RankTotalPoints *int `json:"rankTotalPoints,omitempty"`
		
		RankPercentagePoints *int `json:"rankPercentagePoints,omitempty"`
		Alias
	}{ 
		Points: o.Points,
		
		MaxPoints: o.MaxPoints,
		
		DataPointCount: o.DataPointCount,
		
		PercentOfGoal: o.PercentOfGoal,
		
		RankTotalPoints: o.RankTotalPoints,
		
		RankPercentagePoints: o.RankPercentagePoints,
		Alias:    (Alias)(o),
	})
}

func (o *Insightsdetailsoverallperiodpoints) UnmarshalJSON(b []byte) error {
	var InsightsdetailsoverallperiodpointsMap map[string]interface{}
	err := json.Unmarshal(b, &InsightsdetailsoverallperiodpointsMap)
	if err != nil {
		return err
	}
	
	if Points, ok := InsightsdetailsoverallperiodpointsMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if MaxPoints, ok := InsightsdetailsoverallperiodpointsMap["maxPoints"].(float64); ok {
		MaxPointsInt := int(MaxPoints)
		o.MaxPoints = &MaxPointsInt
	}
	
	if DataPointCount, ok := InsightsdetailsoverallperiodpointsMap["dataPointCount"].(float64); ok {
		DataPointCountInt := int(DataPointCount)
		o.DataPointCount = &DataPointCountInt
	}
	
	if PercentOfGoal, ok := InsightsdetailsoverallperiodpointsMap["percentOfGoal"].(float64); ok {
		o.PercentOfGoal = &PercentOfGoal
	}
    
	if RankTotalPoints, ok := InsightsdetailsoverallperiodpointsMap["rankTotalPoints"].(float64); ok {
		RankTotalPointsInt := int(RankTotalPoints)
		o.RankTotalPoints = &RankTotalPointsInt
	}
	
	if RankPercentagePoints, ok := InsightsdetailsoverallperiodpointsMap["rankPercentagePoints"].(float64); ok {
		RankPercentagePointsInt := int(RankPercentagePoints)
		o.RankPercentagePoints = &RankPercentagePointsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Insightsdetailsoverallperiodpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
