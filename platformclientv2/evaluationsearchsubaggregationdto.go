package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationsearchsubaggregationdto
type Evaluationsearchsubaggregationdto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the aggregation
	Name *string `json:"name,omitempty"`

	// Field - The field to aggregate on.ALLOWED FIELDS BY AGGREGATION TYPE:TERM/COUNT: evaluationStatus, aiScoringFailureType, questionAiAnswerFailureType,TERM: formId, formIdReleased, contextId, questionGroupId, questionId, questionAnswerId, released, questionGroupMarkedNA, questionMarkedNA, questionAiScored, questionEaScored,SUM/AVERAGE/STATS/RANGE: totalScore, totalCriticalScore, questionGroupScorePercentage,criticalQuestionGroupScorePercentage, questionScore,SUM: disputeCount, rescoreCount, eaSuggestionCount, eaAcceptedSuggestionCount,aiSuggestionCount, aiAcceptedSuggestionCount,DATE_HISTOGRAM: conversationDate, createdDate, submittedDate, releaseDate
	Field *string `json:"field,omitempty"`

	// VarType - The Type of Aggregation to Perform
	VarType *string `json:"type,omitempty"`

	// Size - The size limit for term aggregations, 100 size limit for all fields
	Size *int `json:"size,omitempty"`

	// CalendarInterval - The calendar interval for date histogram aggregations
	CalendarInterval *string `json:"calendarInterval,omitempty"`

	// Format - The format for date histogram aggregations
	Format *string `json:"format,omitempty"`

	// Ranges - The ranges for range aggregations
	Ranges *[]Queryapisearchaggregationrange `json:"ranges,omitempty"`

	// SubAggregations - Sub-aggregations to be computed within this aggregation
	SubAggregations *[]Evaluationsearchsubaggregationdto `json:"subAggregations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationsearchsubaggregationdto) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationsearchsubaggregationdto) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationsearchsubaggregationdto
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Field *string `json:"field,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		CalendarInterval *string `json:"calendarInterval,omitempty"`
		
		Format *string `json:"format,omitempty"`
		
		Ranges *[]Queryapisearchaggregationrange `json:"ranges,omitempty"`
		
		SubAggregations *[]Evaluationsearchsubaggregationdto `json:"subAggregations,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Field: o.Field,
		
		VarType: o.VarType,
		
		Size: o.Size,
		
		CalendarInterval: o.CalendarInterval,
		
		Format: o.Format,
		
		Ranges: o.Ranges,
		
		SubAggregations: o.SubAggregations,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationsearchsubaggregationdto) UnmarshalJSON(b []byte) error {
	var EvaluationsearchsubaggregationdtoMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationsearchsubaggregationdtoMap)
	if err != nil {
		return err
	}
	
	if Name, ok := EvaluationsearchsubaggregationdtoMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Field, ok := EvaluationsearchsubaggregationdtoMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if VarType, ok := EvaluationsearchsubaggregationdtoMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Size, ok := EvaluationsearchsubaggregationdtoMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if CalendarInterval, ok := EvaluationsearchsubaggregationdtoMap["calendarInterval"].(string); ok {
		o.CalendarInterval = &CalendarInterval
	}
    
	if Format, ok := EvaluationsearchsubaggregationdtoMap["format"].(string); ok {
		o.Format = &Format
	}
    
	if Ranges, ok := EvaluationsearchsubaggregationdtoMap["ranges"].([]interface{}); ok {
		RangesString, _ := json.Marshal(Ranges)
		json.Unmarshal(RangesString, &o.Ranges)
	}
	
	if SubAggregations, ok := EvaluationsearchsubaggregationdtoMap["subAggregations"].([]interface{}); ok {
		SubAggregationsString, _ := json.Marshal(SubAggregations)
		json.Unmarshal(SubAggregationsString, &o.SubAggregations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationsearchsubaggregationdto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
