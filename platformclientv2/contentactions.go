package platformclientv2
import (
	"encoding/json"
)

// Contentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Contentactions struct { 
	// Url - A URL for a web page to redirect the user to
	Url *string `json:"url,omitempty"`


	// UrlTarget - The target window or tab within the URL's web page. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback - Text to be sent back in reply when a list item is selected
	Textback *string `json:"textback,omitempty"`


	// CommandName - Execute an organization's specific command
	CommandName *string `json:"commandName,omitempty"`


	// Context - Additional context for the command
	Context *map[string]interface{} `json:"context,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentactions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
