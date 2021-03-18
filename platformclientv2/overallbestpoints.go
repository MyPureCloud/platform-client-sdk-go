package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Overallbestpoints
type Overallbestpoints struct { 
	// Division - The requested division
	Division *Division `json:"division,omitempty"`


	// BestPoints - List of gamification best point items
	BestPoints *[]Overallbestpointsitem `json:"bestPoints,omitempty"`

}

// String returns a JSON representation of the model
func (o *Overallbestpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
