package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Columndatatypespecification
type Columndatatypespecification struct { 
	// ColumnName - The column name of a column selected for dynamic queueing
	ColumnName *string `json:"columnName,omitempty"`


	// ColumnDataType - The data type of the column selected for dynamic queueing (TEXT, NUMERIC or TIMESTAMP)
	ColumnDataType *string `json:"columnDataType,omitempty"`


	// Min - The minimum length of the numeric column selected for dynamic queueing
	Min *int `json:"min,omitempty"`


	// Max - The maximum length of the numeric column selected for dynamic queueing
	Max *int `json:"max,omitempty"`


	// MaxLength - The maximum length of the text column selected for dynamic queueing
	MaxLength *int `json:"maxLength,omitempty"`

}

func (o *Columndatatypespecification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Columndatatypespecification
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		ColumnDataType *string `json:"columnDataType,omitempty"`
		
		Min *int `json:"min,omitempty"`
		
		Max *int `json:"max,omitempty"`
		
		MaxLength *int `json:"maxLength,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		ColumnDataType: o.ColumnDataType,
		
		Min: o.Min,
		
		Max: o.Max,
		
		MaxLength: o.MaxLength,
		Alias:    (*Alias)(o),
	})
}

func (o *Columndatatypespecification) UnmarshalJSON(b []byte) error {
	var ColumndatatypespecificationMap map[string]interface{}
	err := json.Unmarshal(b, &ColumndatatypespecificationMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := ColumndatatypespecificationMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if ColumnDataType, ok := ColumndatatypespecificationMap["columnDataType"].(string); ok {
		o.ColumnDataType = &ColumnDataType
	}
    
	if Min, ok := ColumndatatypespecificationMap["min"].(float64); ok {
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := ColumndatatypespecificationMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	
	if MaxLength, ok := ColumndatatypespecificationMap["maxLength"].(float64); ok {
		MaxLengthInt := int(MaxLength)
		o.MaxLength = &MaxLengthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Columndatatypespecification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
