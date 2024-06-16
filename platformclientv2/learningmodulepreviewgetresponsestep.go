package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulepreviewgetresponsestep - Learning module preview get response assignment step
type Learningmodulepreviewgetresponsestep struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of the step
	Id *string `json:"id,omitempty"`

	// ModuleStep - The module step data for this step
	ModuleStep *Learningmoduleinformstep `json:"moduleStep,omitempty"`

	// Structure - The structure for any SCO associated with this step
	Structure *[]Learningmodulepreviewgetscostructure `json:"structure,omitempty"`

	// SuccessStatus - The success status of this step
	SuccessStatus *string `json:"successStatus,omitempty"`

	// CompletionStatus - The completion status of the assignment step
	CompletionStatus *string `json:"completionStatus,omitempty"`

	// CompletionPercentage - The completion percentage for this step
	CompletionPercentage *float32 `json:"completionPercentage,omitempty"`

	// PercentageScore - The percentage score for this step
	PercentageScore *float32 `json:"percentageScore,omitempty"`

	// SignedCookie - The signed cookie information needed to access the content of this step (if required)
	SignedCookie *Learningassignmentstepsignedcookie `json:"signedCookie,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulepreviewgetresponsestep) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulepreviewgetresponsestep) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulepreviewgetresponsestep
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ModuleStep *Learningmoduleinformstep `json:"moduleStep,omitempty"`
		
		Structure *[]Learningmodulepreviewgetscostructure `json:"structure,omitempty"`
		
		SuccessStatus *string `json:"successStatus,omitempty"`
		
		CompletionStatus *string `json:"completionStatus,omitempty"`
		
		CompletionPercentage *float32 `json:"completionPercentage,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		SignedCookie *Learningassignmentstepsignedcookie `json:"signedCookie,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ModuleStep: o.ModuleStep,
		
		Structure: o.Structure,
		
		SuccessStatus: o.SuccessStatus,
		
		CompletionStatus: o.CompletionStatus,
		
		CompletionPercentage: o.CompletionPercentage,
		
		PercentageScore: o.PercentageScore,
		
		SignedCookie: o.SignedCookie,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulepreviewgetresponsestep) UnmarshalJSON(b []byte) error {
	var LearningmodulepreviewgetresponsestepMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulepreviewgetresponsestepMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulepreviewgetresponsestepMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ModuleStep, ok := LearningmodulepreviewgetresponsestepMap["moduleStep"].(map[string]interface{}); ok {
		ModuleStepString, _ := json.Marshal(ModuleStep)
		json.Unmarshal(ModuleStepString, &o.ModuleStep)
	}
	
	if Structure, ok := LearningmodulepreviewgetresponsestepMap["structure"].([]interface{}); ok {
		StructureString, _ := json.Marshal(Structure)
		json.Unmarshal(StructureString, &o.Structure)
	}
	
	if SuccessStatus, ok := LearningmodulepreviewgetresponsestepMap["successStatus"].(string); ok {
		o.SuccessStatus = &SuccessStatus
	}
    
	if CompletionStatus, ok := LearningmodulepreviewgetresponsestepMap["completionStatus"].(string); ok {
		o.CompletionStatus = &CompletionStatus
	}
    
	if CompletionPercentage, ok := LearningmodulepreviewgetresponsestepMap["completionPercentage"].(float64); ok {
		CompletionPercentageFloat32 := float32(CompletionPercentage)
		o.CompletionPercentage = &CompletionPercentageFloat32
	}
	
	if PercentageScore, ok := LearningmodulepreviewgetresponsestepMap["percentageScore"].(float64); ok {
		PercentageScoreFloat32 := float32(PercentageScore)
		o.PercentageScore = &PercentageScoreFloat32
	}
	
	if SignedCookie, ok := LearningmodulepreviewgetresponsestepMap["signedCookie"].(map[string]interface{}); ok {
		SignedCookieString, _ := json.Marshal(SignedCookie)
		json.Unmarshal(SignedCookieString, &o.SignedCookie)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponsestep) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
