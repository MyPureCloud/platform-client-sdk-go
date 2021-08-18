package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleadherencelisting
type Userscheduleadherencelisting struct { 
	// Entities
	Entities *[]Userscheduleadherence `json:"entities,omitempty"`


	// DownloadUrl - The downloadUrl if the response is too large to send directly via http response
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (u *Userscheduleadherencelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userscheduleadherencelisting

	

	return json.Marshal(&struct { 
		Entities *[]Userscheduleadherence `json:"entities,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userscheduleadherencelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
