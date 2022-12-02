package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicobject
type Queueconversationcobrowseeventtopicobject struct { }

func (o *Queueconversationcobrowseeventtopicobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcobrowseeventtopicobject) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicobjectMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
