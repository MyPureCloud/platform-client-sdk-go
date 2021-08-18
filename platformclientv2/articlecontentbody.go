package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Articlecontentbody
type Articlecontentbody struct { 
	// LocationUrl - Presigned URL to retrieve the document content.
	LocationUrl *string `json:"locationUrl,omitempty"`

}

func (u *Articlecontentbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Articlecontentbody

	

	return json.Marshal(&struct { 
		LocationUrl *string `json:"locationUrl,omitempty"`
		*Alias
	}{ 
		LocationUrl: u.LocationUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Articlecontentbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
