package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationform
type Evaluationform struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The evaluation form name
	Name *string `json:"name,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// Published
	Published *bool `json:"published,omitempty"`

	// ContextId
	ContextId *string `json:"contextId,omitempty"`

	// QuestionGroups - A list of question groups
	QuestionGroups *[]Evaluationquestiongroup `json:"questionGroups,omitempty"`

	// PublishedVersions - A list of the published versions of this form. Not populated by default, its availability depends on the endpoint. Use the 'expand=publishHistory' query parameter to retrieve this data where applicable (refer to the endpoint description to see if it is applicable).
	PublishedVersions *Domainentitylistingevaluationform `json:"publishedVersions,omitempty"`

	// EvaluationSettings - Settings for evaluations associated with this form
	EvaluationSettings *Evaluationsettings `json:"evaluationSettings,omitempty"`

	// LatestVersionFormName - The name of the form's most recently published version
	LatestVersionFormName *string `json:"latestVersionFormName,omitempty"`

	// AiScoring - AI scoring settings for the evaluation form.
	AiScoring *Aiscoringsettings `json:"aiScoring,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationform) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationform) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate", }
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
	type Alias Evaluationform
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		QuestionGroups *[]Evaluationquestiongroup `json:"questionGroups,omitempty"`
		
		PublishedVersions *Domainentitylistingevaluationform `json:"publishedVersions,omitempty"`
		
		EvaluationSettings *Evaluationsettings `json:"evaluationSettings,omitempty"`
		
		LatestVersionFormName *string `json:"latestVersionFormName,omitempty"`
		
		AiScoring *Aiscoringsettings `json:"aiScoring,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ModifiedDate: ModifiedDate,
		
		Published: o.Published,
		
		ContextId: o.ContextId,
		
		QuestionGroups: o.QuestionGroups,
		
		PublishedVersions: o.PublishedVersions,
		
		EvaluationSettings: o.EvaluationSettings,
		
		LatestVersionFormName: o.LatestVersionFormName,
		
		AiScoring: o.AiScoring,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationform) UnmarshalJSON(b []byte) error {
	var EvaluationformMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationformMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationformMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluationformMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if modifiedDateString, ok := EvaluationformMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if Published, ok := EvaluationformMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ContextId, ok := EvaluationformMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if QuestionGroups, ok := EvaluationformMap["questionGroups"].([]interface{}); ok {
		QuestionGroupsString, _ := json.Marshal(QuestionGroups)
		json.Unmarshal(QuestionGroupsString, &o.QuestionGroups)
	}
	
	if PublishedVersions, ok := EvaluationformMap["publishedVersions"].(map[string]interface{}); ok {
		PublishedVersionsString, _ := json.Marshal(PublishedVersions)
		json.Unmarshal(PublishedVersionsString, &o.PublishedVersions)
	}
	
	if EvaluationSettings, ok := EvaluationformMap["evaluationSettings"].(map[string]interface{}); ok {
		EvaluationSettingsString, _ := json.Marshal(EvaluationSettings)
		json.Unmarshal(EvaluationSettingsString, &o.EvaluationSettings)
	}
	
	if LatestVersionFormName, ok := EvaluationformMap["latestVersionFormName"].(string); ok {
		o.LatestVersionFormName = &LatestVersionFormName
	}
    
	if AiScoring, ok := EvaluationformMap["aiScoring"].(map[string]interface{}); ok {
		AiScoringString, _ := json.Marshal(AiScoring)
		json.Unmarshal(AiScoringString, &o.AiScoring)
	}
	
	if SelfUri, ok := EvaluationformMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
