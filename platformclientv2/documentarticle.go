package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentarticle
type Documentarticle struct { 
	// Title - The title of the Article.
	Title *string `json:"title,omitempty"`


	// Content - The content of the Article.
	Content *Articlecontent `json:"content,omitempty"`


	// Alternatives - List of Alternative questions related to the title which helps in improving the likelihood of a match to user query.
	Alternatives *[]string `json:"alternatives,omitempty"`

}

func (o *Documentarticle) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentarticle
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Content *Articlecontent `json:"content,omitempty"`
		
		Alternatives *[]string `json:"alternatives,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Content: o.Content,
		
		Alternatives: o.Alternatives,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentarticle) UnmarshalJSON(b []byte) error {
	var DocumentarticleMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentarticleMap)
	if err != nil {
		return err
	}
	
	if Title, ok := DocumentarticleMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Content, ok := DocumentarticleMap["content"].(map[string]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Alternatives, ok := DocumentarticleMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentarticle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
