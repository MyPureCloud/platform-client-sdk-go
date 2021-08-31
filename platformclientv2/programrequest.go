package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programrequest
type Programrequest struct { 
	// Name - The program name
	Name *string `json:"name,omitempty"`


	// Description - The program description
	Description *string `json:"description,omitempty"`


	// TopicIds - The ids of topics associated to the program
	TopicIds *[]string `json:"topicIds,omitempty"`


	// Tags - The program tags
	Tags *[]string `json:"tags,omitempty"`

}

func (o *Programrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		TopicIds *[]string `json:"topicIds,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		TopicIds: o.TopicIds,
		
		Tags: o.Tags,
		Alias:    (*Alias)(o),
	})
}

func (o *Programrequest) UnmarshalJSON(b []byte) error {
	var ProgramrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ProgramrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ProgramrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := ProgramrequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if TopicIds, ok := ProgramrequestMap["topicIds"].([]interface{}); ok {
		TopicIdsString, _ := json.Marshal(TopicIds)
		json.Unmarshal(TopicIdsString, &o.TopicIds)
	}
	
	if Tags, ok := ProgramrequestMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Programrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
