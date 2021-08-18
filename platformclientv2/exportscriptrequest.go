package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Exportscriptrequest - Creating an exported script via Download Service
type Exportscriptrequest struct { 
	// FileName - The final file name (no extension) of the script download: <fileName>.script
	FileName *string `json:"fileName,omitempty"`


	// VersionId - The UUID version of the script to be exported.  Defaults to the current editable version.
	VersionId *string `json:"versionId,omitempty"`

}

func (u *Exportscriptrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Exportscriptrequest

	

	return json.Marshal(&struct { 
		FileName *string `json:"fileName,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		*Alias
	}{ 
		FileName: u.FileName,
		
		VersionId: u.VersionId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Exportscriptrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
