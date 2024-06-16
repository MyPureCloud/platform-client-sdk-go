package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulepreviewupdatescostructure - Learning module preview update SCO structure
type Learningmodulepreviewupdatescostructure struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of this SCO in the course manifest
	Id *string `json:"id,omitempty"`

	// Name - The name of this SCO in the course manifest
	Name *string `json:"name,omitempty"`

	// SuccessStatus - The success status of this SCO
	SuccessStatus *string `json:"successStatus,omitempty"`

	// CompletionStatus - The completion status of this SCO
	CompletionStatus *string `json:"completionStatus,omitempty"`

	// PercentageScore - Percentage Score
	PercentageScore *float32 `json:"percentageScore,omitempty"`

	// Children - Child items belonging to this SCO in the course manifest
	Children *[]Learningmodulepreviewupdatescostructure `json:"children,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulepreviewupdatescostructure) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulepreviewupdatescostructure) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulepreviewupdatescostructure
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SuccessStatus *string `json:"successStatus,omitempty"`
		
		CompletionStatus *string `json:"completionStatus,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		Children *[]Learningmodulepreviewupdatescostructure `json:"children,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SuccessStatus: o.SuccessStatus,
		
		CompletionStatus: o.CompletionStatus,
		
		PercentageScore: o.PercentageScore,
		
		Children: o.Children,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulepreviewupdatescostructure) UnmarshalJSON(b []byte) error {
	var LearningmodulepreviewupdatescostructureMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulepreviewupdatescostructureMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulepreviewupdatescostructureMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LearningmodulepreviewupdatescostructureMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SuccessStatus, ok := LearningmodulepreviewupdatescostructureMap["successStatus"].(string); ok {
		o.SuccessStatus = &SuccessStatus
	}
    
	if CompletionStatus, ok := LearningmodulepreviewupdatescostructureMap["completionStatus"].(string); ok {
		o.CompletionStatus = &CompletionStatus
	}
    
	if PercentageScore, ok := LearningmodulepreviewupdatescostructureMap["percentageScore"].(float64); ok {
		PercentageScoreFloat32 := float32(PercentageScore)
		o.PercentageScore = &PercentageScoreFloat32
	}
	
	if Children, ok := LearningmodulepreviewupdatescostructureMap["children"].([]interface{}); ok {
		ChildrenString, _ := json.Marshal(Children)
		json.Unmarshal(ChildrenString, &o.Children)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdatescostructure) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
