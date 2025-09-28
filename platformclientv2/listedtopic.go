package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Listedtopic
type Listedtopic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Published
	Published *bool `json:"published,omitempty"`

	// Strictness
	Strictness *string `json:"strictness,omitempty"`

	// MatchingType
	MatchingType *string `json:"matchingType,omitempty"`

	// ProgramsCount
	ProgramsCount *int `json:"programsCount,omitempty"`

	// Tags
	Tags *[]string `json:"tags,omitempty"`

	// Dialect
	Dialect *string `json:"dialect,omitempty"`

	// Participants
	Participants *string `json:"participants,omitempty"`

	// PhrasesCount
	PhrasesCount *int `json:"phrasesCount,omitempty"`

	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`

	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Listedtopic) SetField(field string, fieldValue interface{}) {
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

func (o Listedtopic) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified", }
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
	type Alias Listedtopic
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		MatchingType *string `json:"matchingType,omitempty"`
		
		ProgramsCount *int `json:"programsCount,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		PhrasesCount *int `json:"phrasesCount,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Published: o.Published,
		
		Strictness: o.Strictness,
		
		MatchingType: o.MatchingType,
		
		ProgramsCount: o.ProgramsCount,
		
		Tags: o.Tags,
		
		Dialect: o.Dialect,
		
		Participants: o.Participants,
		
		PhrasesCount: o.PhrasesCount,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Listedtopic) UnmarshalJSON(b []byte) error {
	var ListedtopicMap map[string]interface{}
	err := json.Unmarshal(b, &ListedtopicMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ListedtopicMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ListedtopicMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ListedtopicMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := ListedtopicMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if Strictness, ok := ListedtopicMap["strictness"].(string); ok {
		o.Strictness = &Strictness
	}
    
	if MatchingType, ok := ListedtopicMap["matchingType"].(string); ok {
		o.MatchingType = &MatchingType
	}
    
	if ProgramsCount, ok := ListedtopicMap["programsCount"].(float64); ok {
		ProgramsCountInt := int(ProgramsCount)
		o.ProgramsCount = &ProgramsCountInt
	}
	
	if Tags, ok := ListedtopicMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Dialect, ok := ListedtopicMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
    
	if Participants, ok := ListedtopicMap["participants"].(string); ok {
		o.Participants = &Participants
	}
    
	if PhrasesCount, ok := ListedtopicMap["phrasesCount"].(float64); ok {
		PhrasesCountInt := int(PhrasesCount)
		o.PhrasesCount = &PhrasesCountInt
	}
	
	if ModifiedBy, ok := ListedtopicMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := ListedtopicMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := ListedtopicMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Listedtopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
