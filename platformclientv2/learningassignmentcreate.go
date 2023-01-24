package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentcreate
type Learningassignmentcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ModuleId - The Learning module Id associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`

	// UserId - The User for whom the assignment is assigned
	UserId *string `json:"userId,omitempty"`

	// RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecommendedCompletionDate *time.Time `json:"recommendedCompletionDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassignmentcreate) SetField(field string, fieldValue interface{}) {
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

func (o Learningassignmentcreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "RecommendedCompletionDate", }
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
	type Alias Learningassignmentcreate
	
	RecommendedCompletionDate := new(string)
	if o.RecommendedCompletionDate != nil {
		
		*RecommendedCompletionDate = timeutil.Strftime(o.RecommendedCompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecommendedCompletionDate = nil
	}
	
	return json.Marshal(&struct { 
		ModuleId *string `json:"moduleId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		RecommendedCompletionDate *string `json:"recommendedCompletionDate,omitempty"`
		Alias
	}{ 
		ModuleId: o.ModuleId,
		
		UserId: o.UserId,
		
		RecommendedCompletionDate: RecommendedCompletionDate,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassignmentcreate) UnmarshalJSON(b []byte) error {
	var LearningassignmentcreateMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentcreateMap)
	if err != nil {
		return err
	}
	
	if ModuleId, ok := LearningassignmentcreateMap["moduleId"].(string); ok {
		o.ModuleId = &ModuleId
	}
    
	if UserId, ok := LearningassignmentcreateMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if recommendedCompletionDateString, ok := LearningassignmentcreateMap["recommendedCompletionDate"].(string); ok {
		RecommendedCompletionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", recommendedCompletionDateString)
		o.RecommendedCompletionDate = &RecommendedCompletionDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
