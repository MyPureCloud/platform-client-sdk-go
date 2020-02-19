package platformclientv2
import (
	"encoding/json"
)

// Gkndocumentationresult
type Gkndocumentationresult struct { 
	// Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content *string `json:"content,omitempty"`


	// Link - URL link for the documentation entity. Will be returned in responses for certain entities.
	Link *string `json:"link,omitempty"`


	// Title - The title of the documentation entity. Will be returned in responses for certain entities.
	Title *string `json:"title,omitempty"`


	// VarType - The search type. Will be returned in responses for certain entities.
	VarType *string `json:"_type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Gkndocumentationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
