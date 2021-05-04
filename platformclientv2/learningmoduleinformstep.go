package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmoduleinformstep
type Learningmoduleinformstep struct { 
	// VarType - The learning module inform step type
	VarType *string `json:"type,omitempty"`


	// Name - The name of the inform step or content
	Name *string `json:"name,omitempty"`


	// Value - The value for inform step
	Value *string `json:"value,omitempty"`


	// SharingUri - The sharing uri for Content type inform step
	SharingUri *string `json:"sharingUri,omitempty"`


	// ContentType - The document type for Content type Inform step
	ContentType *string `json:"contentType,omitempty"`


	// Order - The order of inform step in a learning module
	Order *int `json:"order,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningmoduleinformstep) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
