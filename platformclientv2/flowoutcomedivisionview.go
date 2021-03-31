package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Flowoutcomedivisionview
type Flowoutcomedivisionview struct { 
	// Id - The flow outcome identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow outcome name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowoutcomedivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
