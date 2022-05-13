package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversion
type Nludomainversion struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Domain - The NLU domain of the version.
	Domain **Nludomain `json:"domain,omitempty"`


	// Description - The description of the NLU domain version.
	Description *string `json:"description,omitempty"`


	// Language - The language that the NLU domain version supports.
	Language *string `json:"language,omitempty"`


	// Published - Whether this NLU domain version has been published.
	Published *bool `json:"published,omitempty"`


	// DateCreated - The date when the NLU domain version was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date when the NLU domain version was updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateTrained - The date when the NLU domain version was trained. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTrained *time.Time `json:"dateTrained,omitempty"`


	// DatePublished - The date when the NLU domain version was published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// TrainingStatus - The training status of the NLU domain version.
	TrainingStatus *string `json:"trainingStatus,omitempty"`


	// EvaluationStatus - The evaluation status of the NLU domain version.
	EvaluationStatus *string `json:"evaluationStatus,omitempty"`


	// Intents - The intents defined for this NLU domain version.
	Intents *[]Intentdefinition `json:"intents,omitempty"`


	// EntityTypes - The entity types defined for this NLU domain version.
	EntityTypes *[]Namedentitytypedefinition `json:"entityTypes,omitempty"`


	// Entities - The entities defined for this NLU domain version.This field is mutually exclusive with entityTypeBindings
	Entities *[]Namedentitydefinition `json:"entities,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Nludomainversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludomainversion
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateTrained := new(string)
	if o.DateTrained != nil {
		
		*DateTrained = timeutil.Strftime(o.DateTrained, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTrained = nil
	}
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Domain **Nludomain `json:"domain,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateTrained *string `json:"dateTrained,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		TrainingStatus *string `json:"trainingStatus,omitempty"`
		
		EvaluationStatus *string `json:"evaluationStatus,omitempty"`
		
		Intents *[]Intentdefinition `json:"intents,omitempty"`
		
		EntityTypes *[]Namedentitytypedefinition `json:"entityTypes,omitempty"`
		
		Entities *[]Namedentitydefinition `json:"entities,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Domain: o.Domain,
		
		Description: o.Description,
		
		Language: o.Language,
		
		Published: o.Published,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateTrained: DateTrained,
		
		DatePublished: DatePublished,
		
		TrainingStatus: o.TrainingStatus,
		
		EvaluationStatus: o.EvaluationStatus,
		
		Intents: o.Intents,
		
		EntityTypes: o.EntityTypes,
		
		Entities: o.Entities,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludomainversion) UnmarshalJSON(b []byte) error {
	var NludomainversionMap map[string]interface{}
	err := json.Unmarshal(b, &NludomainversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NludomainversionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Domain, ok := NludomainversionMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	
	if Description, ok := NludomainversionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Language, ok := NludomainversionMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Published, ok := NludomainversionMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if dateCreatedString, ok := NludomainversionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := NludomainversionMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateTrainedString, ok := NludomainversionMap["dateTrained"].(string); ok {
		DateTrained, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateTrainedString)
		o.DateTrained = &DateTrained
	}
	
	if datePublishedString, ok := NludomainversionMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if TrainingStatus, ok := NludomainversionMap["trainingStatus"].(string); ok {
		o.TrainingStatus = &TrainingStatus
	}
    
	if EvaluationStatus, ok := NludomainversionMap["evaluationStatus"].(string); ok {
		o.EvaluationStatus = &EvaluationStatus
	}
    
	if Intents, ok := NludomainversionMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if EntityTypes, ok := NludomainversionMap["entityTypes"].([]interface{}); ok {
		EntityTypesString, _ := json.Marshal(EntityTypes)
		json.Unmarshal(EntityTypesString, &o.EntityTypes)
	}
	
	if Entities, ok := NludomainversionMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if SelfUri, ok := NludomainversionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nludomainversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
