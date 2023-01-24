package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapestimateresult
type Actionmapestimateresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QualifiedSessionCount - Number of sessions qualified for Action map.
	QualifiedSessionCount *int `json:"qualifiedSessionCount,omitempty"`

	// TotalSessionCount - Total number of sessions.
	TotalSessionCount *int `json:"totalSessionCount,omitempty"`

	// PerSegmentCounts - Number of sessions qualified for Action map per segment.
	PerSegmentCounts *[]Segmentestimatecount `json:"perSegmentCounts,omitempty"`

	// OutcomesScoresCount - Difference made by outcome criteria to number of sessions qualified for Action map.
	OutcomesScoresCount *int `json:"outcomesScoresCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionmapestimateresult) SetField(field string, fieldValue interface{}) {
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

func (o Actionmapestimateresult) MarshalJSON() ([]byte, error) {
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
	type Alias Actionmapestimateresult
	
	return json.Marshal(&struct { 
		QualifiedSessionCount *int `json:"qualifiedSessionCount,omitempty"`
		
		TotalSessionCount *int `json:"totalSessionCount,omitempty"`
		
		PerSegmentCounts *[]Segmentestimatecount `json:"perSegmentCounts,omitempty"`
		
		OutcomesScoresCount *int `json:"outcomesScoresCount,omitempty"`
		Alias
	}{ 
		QualifiedSessionCount: o.QualifiedSessionCount,
		
		TotalSessionCount: o.TotalSessionCount,
		
		PerSegmentCounts: o.PerSegmentCounts,
		
		OutcomesScoresCount: o.OutcomesScoresCount,
		Alias:    (Alias)(o),
	})
}

func (o *Actionmapestimateresult) UnmarshalJSON(b []byte) error {
	var ActionmapestimateresultMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapestimateresultMap)
	if err != nil {
		return err
	}
	
	if QualifiedSessionCount, ok := ActionmapestimateresultMap["qualifiedSessionCount"].(float64); ok {
		QualifiedSessionCountInt := int(QualifiedSessionCount)
		o.QualifiedSessionCount = &QualifiedSessionCountInt
	}
	
	if TotalSessionCount, ok := ActionmapestimateresultMap["totalSessionCount"].(float64); ok {
		TotalSessionCountInt := int(TotalSessionCount)
		o.TotalSessionCount = &TotalSessionCountInt
	}
	
	if PerSegmentCounts, ok := ActionmapestimateresultMap["perSegmentCounts"].([]interface{}); ok {
		PerSegmentCountsString, _ := json.Marshal(PerSegmentCounts)
		json.Unmarshal(PerSegmentCountsString, &o.PerSegmentCounts)
	}
	
	if OutcomesScoresCount, ok := ActionmapestimateresultMap["outcomesScoresCount"].(float64); ok {
		OutcomesScoresCountInt := int(OutcomesScoresCount)
		o.OutcomesScoresCount = &OutcomesScoresCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapestimateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
