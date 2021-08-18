package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Authenticationsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authenticationsettings

	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		
		IntegrationId: u.IntegrationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Authenticationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
