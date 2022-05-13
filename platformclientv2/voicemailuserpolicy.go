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

func (o *Voicemailuserpolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailuserpolicy
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Enabled: o.Enabled,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		Pin: o.Pin,
		
		ModifiedDate: ModifiedDate,
		
		SendEmailNotifications: o.SendEmailNotifications,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailuserpolicy) UnmarshalJSON(b []byte) error {
	var VoicemailuserpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailuserpolicyMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := VoicemailuserpolicyMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if AlertTimeoutSeconds, ok := VoicemailuserpolicyMap["alertTimeoutSeconds"].(float64); ok {
		AlertTimeoutSecondsInt := int(AlertTimeoutSeconds)
		o.AlertTimeoutSeconds = &AlertTimeoutSecondsInt
	}
	
	if Pin, ok := VoicemailuserpolicyMap["pin"].(string); ok {
		o.Pin = &Pin
	}
    
	if modifiedDateString, ok := VoicemailuserpolicyMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SendEmailNotifications, ok := VoicemailuserpolicyMap["sendEmailNotifications"].(bool); ok {
		o.SendEmailNotifications = &SendEmailNotifications
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailuserpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
