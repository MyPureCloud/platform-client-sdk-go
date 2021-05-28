package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Contentactions struct { 
	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`


	// UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback - Text to be sent back in reply when the item is selected.
	Textback *string `json:"textback,omitempty"`


	// CommandName - Execute an organization's specific command.
	CommandName *string `json:"commandName,omitempty"`


	// Context - Additional context for the command.
	Context *map[string]interface{} `json:"context,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
