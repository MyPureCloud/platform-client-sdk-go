package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningcoverartthumbnail
type Learningcoverartthumbnail struct { 
	// Resolution - Resolution of thumbnail
	Resolution *string `json:"resolution,omitempty"`


	// Url - The URL for the thumbnail
	Url *string `json:"url,omitempty"`

}

func (o *Learningcoverartthumbnail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningcoverartthumbnail
	
	return json.Marshal(&struct { 
		Resolution *string `json:"resolution,omitempty"`
		
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Resolution: o.Resolution,
		
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningcoverartthumbnail) UnmarshalJSON(b []byte) error {
	var LearningcoverartthumbnailMap map[string]interface{}
	err := json.Unmarshal(b, &LearningcoverartthumbnailMap)
	if err != nil {
		return err
	}
	
	if Resolution, ok := LearningcoverartthumbnailMap["resolution"].(string); ok {
		o.Resolution = &Resolution
	}
    
	if Url, ok := LearningcoverartthumbnailMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningcoverartthumbnail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
