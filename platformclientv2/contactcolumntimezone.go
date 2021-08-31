package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumntimezone
type Contactcolumntimezone struct { 
	// TimeZone - Time zone that the column matched to. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`


	// ColumnType - Column Type will be either PHONE or ZIP
	ColumnType *string `json:"columnType,omitempty"`

}

func (o *Contactcolumntimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcolumntimezone
	
	return json.Marshal(&struct { 
		TimeZone *string `json:"timeZone,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		*Alias
	}{ 
		TimeZone: o.TimeZone,
		
		ColumnType: o.ColumnType,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactcolumntimezone) UnmarshalJSON(b []byte) error {
	var ContactcolumntimezoneMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcolumntimezoneMap)
	if err != nil {
		return err
	}
	
	if TimeZone, ok := ContactcolumntimezoneMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if ColumnType, ok := ContactcolumntimezoneMap["columnType"].(string); ok {
		o.ColumnType = &ColumnType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcolumntimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
