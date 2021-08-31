package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topic
type Topic struct { 
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


	// Programs
	Programs *[]Baseprogramentity `json:"programs,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// Dialect
	Dialect *string `json:"dialect,omitempty"`


	// Participants
	Participants *string `json:"participants,omitempty"`


	// Phrases
	Phrases *[]Phrase `json:"phrases,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// PublishedBy
	PublishedBy *Addressableentityref `json:"publishedBy,omitempty"`


	// DatePublished - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Topic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topic
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		Programs *[]Baseprogramentity `json:"programs,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		Phrases *[]Phrase `json:"phrases,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		PublishedBy *Addressableentityref `json:"publishedBy,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Published: o.Published,
		
		Strictness: o.Strictness,
		
		Programs: o.Programs,
		
		Tags: o.Tags,
		
		Dialect: o.Dialect,
		
		Participants: o.Participants,
		
		Phrases: o.Phrases,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		PublishedBy: o.PublishedBy,
		
		DatePublished: DatePublished,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Topic) UnmarshalJSON(b []byte) error {
	var TopicMap map[string]interface{}
	err := json.Unmarshal(b, &TopicMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TopicMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TopicMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := TopicMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Published, ok := TopicMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if Strictness, ok := TopicMap["strictness"].(string); ok {
		o.Strictness = &Strictness
	}
	
	if Programs, ok := TopicMap["programs"].([]interface{}); ok {
		ProgramsString, _ := json.Marshal(Programs)
		json.Unmarshal(ProgramsString, &o.Programs)
	}
	
	if Tags, ok := TopicMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Dialect, ok := TopicMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
	
	if Participants, ok := TopicMap["participants"].(string); ok {
		o.Participants = &Participants
	}
	
	if Phrases, ok := TopicMap["phrases"].([]interface{}); ok {
		PhrasesString, _ := json.Marshal(Phrases)
		json.Unmarshal(PhrasesString, &o.Phrases)
	}
	
	if ModifiedBy, ok := TopicMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := TopicMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if PublishedBy, ok := TopicMap["publishedBy"].(map[string]interface{}); ok {
		PublishedByString, _ := json.Marshal(PublishedBy)
		json.Unmarshal(PublishedByString, &o.PublishedBy)
	}
	
	if datePublishedString, ok := TopicMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if SelfUri, ok := TopicMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
