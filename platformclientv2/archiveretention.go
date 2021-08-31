package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Archiveretention
type Archiveretention struct { 
	// Days
	Days *int `json:"days,omitempty"`


	// StorageMedium
	StorageMedium *string `json:"storageMedium,omitempty"`

}

func (o *Archiveretention) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Archiveretention
	
	return json.Marshal(&struct { 
		Days *int `json:"days,omitempty"`
		
		StorageMedium *string `json:"storageMedium,omitempty"`
		*Alias
	}{ 
		Days: o.Days,
		
		StorageMedium: o.StorageMedium,
		Alias:    (*Alias)(o),
	})
}

func (o *Archiveretention) UnmarshalJSON(b []byte) error {
	var ArchiveretentionMap map[string]interface{}
	err := json.Unmarshal(b, &ArchiveretentionMap)
	if err != nil {
		return err
	}
	
	if Days, ok := ArchiveretentionMap["days"].(float64); ok {
		DaysInt := int(Days)
		o.Days = &DaysInt
	}
	
	if StorageMedium, ok := ArchiveretentionMap["storageMedium"].(string); ok {
		o.StorageMedium = &StorageMedium
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Archiveretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
