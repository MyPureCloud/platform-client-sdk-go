package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Documentarticle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
