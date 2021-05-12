package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Evententity
type Evententity struct { 
	// EntityType - Type of entity the event pertains to. e.g. integration
	EntityType *string `json:"entityType,omitempty"`


	// Id - ID of the entity the event pertains to.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
