package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listedprogram
type Listedprogram struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// TopicsCount
	TopicsCount *int `json:"topicsCount,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Listedprogram) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listedprogram
	
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
		
		TopicsCount *int `json:"topicsCount,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Published: o.Published,
		
		TopicsCount: o.TopicsCount,
		
		Tags: o.Tags,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Listedprogram) UnmarshalJSON(b []byte) error {
	var ListedprogramMap map[string]interface{}
	err := json.Unmarshal(b, &ListedprogramMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ListedprogramMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ListedprogramMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ListedprogramMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := ListedprogramMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if TopicsCount, ok := ListedprogramMap["topicsCount"].(float64); ok {
		TopicsCountInt := int(TopicsCount)
		o.TopicsCount = &TopicsCountInt
	}
	
	if Tags, ok := ListedprogramMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if ModifiedBy, ok := ListedprogramMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := ListedprogramMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := ListedprogramMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Listedprogram) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
