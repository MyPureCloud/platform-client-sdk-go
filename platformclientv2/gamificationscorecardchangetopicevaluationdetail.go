package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicevaluationdetail
type Gamificationscorecardchangetopicevaluationdetail struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EvaluationId
	EvaluationId *string `json:"evaluationId,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// ConversationDate
	ConversationDate *string `json:"conversationDate,omitempty"`

	// FormName
	FormName *string `json:"formName,omitempty"`

	// Points
	Points *int `json:"points,omitempty"`

	// MaxPoints
	MaxPoints *int `json:"maxPoints,omitempty"`

	// EvaluationScore
	EvaluationScore *int `json:"evaluationScore,omitempty"`

	// EvaluationScoreDouble
	EvaluationScoreDouble *float32 `json:"evaluationScoreDouble,omitempty"`

	// MediaTypes
	MediaTypes *[]string `json:"mediaTypes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gamificationscorecardchangetopicevaluationdetail) SetField(field string, fieldValue interface{}) {
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

func (o Gamificationscorecardchangetopicevaluationdetail) MarshalJSON() ([]byte, error) {
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
	type Alias Gamificationscorecardchangetopicevaluationdetail
	
	return json.Marshal(&struct { 
		EvaluationId *string `json:"evaluationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		FormName *string `json:"formName,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		EvaluationScore *int `json:"evaluationScore,omitempty"`
		
		EvaluationScoreDouble *float32 `json:"evaluationScoreDouble,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		Alias
	}{ 
		EvaluationId: o.EvaluationId,
		
		ConversationId: o.ConversationId,
		
		ConversationDate: o.ConversationDate,
		
		FormName: o.FormName,
		
		Points: o.Points,
		
		MaxPoints: o.MaxPoints,
		
		EvaluationScore: o.EvaluationScore,
		
		EvaluationScoreDouble: o.EvaluationScoreDouble,
		
		MediaTypes: o.MediaTypes,
		Alias:    (Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicevaluationdetail) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicevaluationdetailMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicevaluationdetailMap)
	if err != nil {
		return err
	}
	
	if EvaluationId, ok := GamificationscorecardchangetopicevaluationdetailMap["evaluationId"].(string); ok {
		o.EvaluationId = &EvaluationId
	}
    
	if ConversationId, ok := GamificationscorecardchangetopicevaluationdetailMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ConversationDate, ok := GamificationscorecardchangetopicevaluationdetailMap["conversationDate"].(string); ok {
		o.ConversationDate = &ConversationDate
	}
    
	if FormName, ok := GamificationscorecardchangetopicevaluationdetailMap["formName"].(string); ok {
		o.FormName = &FormName
	}
    
	if Points, ok := GamificationscorecardchangetopicevaluationdetailMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if MaxPoints, ok := GamificationscorecardchangetopicevaluationdetailMap["maxPoints"].(float64); ok {
		MaxPointsInt := int(MaxPoints)
		o.MaxPoints = &MaxPointsInt
	}
	
	if EvaluationScore, ok := GamificationscorecardchangetopicevaluationdetailMap["evaluationScore"].(float64); ok {
		EvaluationScoreInt := int(EvaluationScore)
		o.EvaluationScore = &EvaluationScoreInt
	}
	
	if EvaluationScoreDouble, ok := GamificationscorecardchangetopicevaluationdetailMap["evaluationScoreDouble"].(float64); ok {
		EvaluationScoreDoubleFloat32 := float32(EvaluationScoreDouble)
		o.EvaluationScoreDouble = &EvaluationScoreDoubleFloat32
	}
    
	if MediaTypes, ok := GamificationscorecardchangetopicevaluationdetailMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicevaluationdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
