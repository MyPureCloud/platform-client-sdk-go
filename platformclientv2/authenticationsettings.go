package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Authenticationsettings - Settings for authenticated webdeployments.
type Authenticationsettings struct { 
	// Enabled - Indicate if these auth is required for this deployment. If, for example, this flag is set to true then webmessaging sessions can not send messages unless the end-user is authenticated.
	Enabled *bool `json:"enabled,omitempty"`


	// IntegrationId - The integration identifier which contains the auth settings required on the deployment.
	IntegrationId *string `json:"integrationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authenticationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
