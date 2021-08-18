package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailuserpolicy
type Voicemailuserpolicy struct { 
	// Enabled - Whether the user has voicemail enabled
	Enabled *bool `json:"enabled,omitempty"`


	// AlertTimeoutSeconds - The number of seconds to ring the user's phone before a call is transfered to voicemail
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`


	// Pin - The user's PIN to access their voicemail. This property is only used for updates and never provided otherwise to ensure security
	Pin *string `json:"pin,omitempty"`


	// ModifiedDate - The date the policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SendEmailNotifications - Whether email notifications are sent to the user when a new voicemail is received
	SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`

}

func (u *Voicemailuserpolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailuserpolicy

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		Pin *string `json:"pin,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		
		AlertTimeoutSeconds: u.AlertTimeoutSeconds,
		
		Pin: u.Pin,
		
		ModifiedDate: ModifiedDate,
		
		SendEmailNotifications: u.SendEmailNotifications,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailuserpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
