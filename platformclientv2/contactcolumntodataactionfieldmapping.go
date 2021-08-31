package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { }

func (o *Contactcolumntodataactionfieldmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcolumntodataactionfieldmapping
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Contactcolumntodataactionfieldmapping) UnmarshalJSON(b []byte) error {
	var ContactcolumntodataactionfieldmappingMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcolumntodataactionfieldmappingMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
