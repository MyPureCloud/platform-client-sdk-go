package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listedtopic
type Listedtopic struct { 
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

func (o *Listedtopic) MarshalJSON() ([]byte, error) {
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
		
		ProgramsCount *int `json:"programsCount,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		PhrasesCount *int `json:"phrasesCount,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Published: o.Published,
		
		Strictness: o.Strictness,
		
		ProgramsCount: o.ProgramsCount,
		
		Tags: o.Tags,
		
		Dialect: o.Dialect,
		
		Participants: o.Participants,
		
		PhrasesCount: o.PhrasesCount,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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
