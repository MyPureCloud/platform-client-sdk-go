package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityevaluationscoreitem
type Qualityevaluationscoreitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EvaluationId - The id of evaluation
	EvaluationId *string `json:"evaluationId,omitempty"`

	// ConversationId - The id of conversation
	ConversationId *string `json:"conversationId,omitempty"`

	// ConversationDate - The date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationDate *time.Time `json:"conversationDate,omitempty"`

	// ConversationEndDate - The end date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationEndDate *time.Time `json:"conversationEndDate,omitempty"`

	// FormName - The name of form
	FormName *string `json:"formName,omitempty"`

	// Points - Gamification points earned for this metric
	Points *int `json:"points,omitempty"`

	// EvaluationScore - The quality score of evaluation as a percentage
	EvaluationScore *float64 `json:"evaluationScore,omitempty"`

	// MaxPoints - Gamification max points for this metric
	MaxPoints *int `json:"maxPoints,omitempty"`

	// MediaTypes - A list of media types for the metric
	MediaTypes *[]string `json:"mediaTypes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Qualityevaluationscoreitem) SetField(field string, fieldValue interface{}) {
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

func (o Qualityevaluationscoreitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ConversationDate","ConversationEndDate", }
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
	type Alias Qualityevaluationscoreitem
	
	ConversationDate := new(string)
	if o.ConversationDate != nil {
		
		*ConversationDate = timeutil.Strftime(o.ConversationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationDate = nil
	}
	
	ConversationEndDate := new(string)
	if o.ConversationEndDate != nil {
		
		*ConversationEndDate = timeutil.Strftime(o.ConversationEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEndDate = nil
	}
	
	return json.Marshal(&struct { 
		EvaluationId *string `json:"evaluationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		ConversationEndDate *string `json:"conversationEndDate,omitempty"`
		
		FormName *string `json:"formName,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		EvaluationScore *float64 `json:"evaluationScore,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		Alias
	}{ 
		EvaluationId: o.EvaluationId,
		
		ConversationId: o.ConversationId,
		
		ConversationDate: ConversationDate,
		
		ConversationEndDate: ConversationEndDate,
		
		FormName: o.FormName,
		
		Points: o.Points,
		
		EvaluationScore: o.EvaluationScore,
		
		MaxPoints: o.MaxPoints,
		
		MediaTypes: o.MediaTypes,
		Alias:    (Alias)(o),
	})
}

func (o *Qualityevaluationscoreitem) UnmarshalJSON(b []byte) error {
	var QualityevaluationscoreitemMap map[string]interface{}
	err := json.Unmarshal(b, &QualityevaluationscoreitemMap)
	if err != nil {
		return err
	}
	
	if EvaluationId, ok := QualityevaluationscoreitemMap["evaluationId"].(string); ok {
		o.EvaluationId = &EvaluationId
	}
    
	if ConversationId, ok := QualityevaluationscoreitemMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if conversationDateString, ok := QualityevaluationscoreitemMap["conversationDate"].(string); ok {
		ConversationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationDateString)
		o.ConversationDate = &ConversationDate
	}
	
	if conversationEndDateString, ok := QualityevaluationscoreitemMap["conversationEndDate"].(string); ok {
		ConversationEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationEndDateString)
		o.ConversationEndDate = &ConversationEndDate
	}
	
	if FormName, ok := QualityevaluationscoreitemMap["formName"].(string); ok {
		o.FormName = &FormName
	}
    
	if Points, ok := QualityevaluationscoreitemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if EvaluationScore, ok := QualityevaluationscoreitemMap["evaluationScore"].(float64); ok {
		o.EvaluationScore = &EvaluationScore
	}
    
	if MaxPoints, ok := QualityevaluationscoreitemMap["maxPoints"].(float64); ok {
		MaxPointsInt := int(MaxPoints)
		o.MaxPoints = &MaxPointsInt
	}
	
	if MediaTypes, ok := QualityevaluationscoreitemMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Qualityevaluationscoreitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
