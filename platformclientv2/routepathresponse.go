package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathresponse
type Routepathresponse struct { 
	// Queue - The ID of the queue associated with the route path
	Queue *Queuereference `json:"queue,omitempty"`


	// MediaType - The media type of the given queue associated with the route path
	MediaType *string `json:"mediaType,omitempty"`


	// Language - The ID of the language associated with the route path
	Language *Languagereference `json:"language,omitempty"`


	// Skills - The set of skills associated with the route path
	Skills *[]Routingskillreference `json:"skills,omitempty"`

}

func (o *Routepathresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routepathresponse
	
	return json.Marshal(&struct { 
		Queue *Queuereference `json:"queue,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Language *Languagereference `json:"language,omitempty"`
		
		Skills *[]Routingskillreference `json:"skills,omitempty"`
		*Alias
	}{ 
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		Language: o.Language,
		
		Skills: o.Skills,
		Alias:    (*Alias)(o),
	})
}

func (o *Routepathresponse) UnmarshalJSON(b []byte) error {
	var RoutepathresponseMap map[string]interface{}
	err := json.Unmarshal(b, &RoutepathresponseMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := RoutepathresponseMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := RoutepathresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Language, ok := RoutepathresponseMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Skills, ok := RoutepathresponseMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routepathresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
