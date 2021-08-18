package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathresponse - Route path configuration
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

func (u *Routepathresponse) MarshalJSON() ([]byte, error) {
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
		Queue: u.Queue,
		
		MediaType: u.MediaType,
		
		Language: u.Language,
		
		Skills: u.Skills,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routepathresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
