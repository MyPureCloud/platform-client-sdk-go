package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Program
type Program struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// Topics
	Topics *[]Basetopicentitiy `json:"topics,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


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

func (o *Program) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Program
	
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
		
		Topics *[]Basetopicentitiy `json:"topics,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
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
		
		Topics: o.Topics,
		
		Tags: o.Tags,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		PublishedBy: o.PublishedBy,
		
		DatePublished: DatePublished,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Program) UnmarshalJSON(b []byte) error {
	var ProgramMap map[string]interface{}
	err := json.Unmarshal(b, &ProgramMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ProgramMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ProgramMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ProgramMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := ProgramMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if Topics, ok := ProgramMap["topics"].([]interface{}); ok {
		TopicsString, _ := json.Marshal(Topics)
		json.Unmarshal(TopicsString, &o.Topics)
	}
	
	if Tags, ok := ProgramMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if ModifiedBy, ok := ProgramMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := ProgramMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if PublishedBy, ok := ProgramMap["publishedBy"].(map[string]interface{}); ok {
		PublishedByString, _ := json.Marshal(PublishedBy)
		json.Unmarshal(PublishedByString, &o.PublishedBy)
	}
	
	if datePublishedString, ok := ProgramMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if SelfUri, ok := ProgramMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Program) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
