package types

import (
	"database/sql/driver"

	"encoding/json"
)

type SQLJSONMODEL struct{}

// Message Model JSon Parse on db
func (b SQLJSONMODEL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), &b)
}
func (b SQLJSONMODEL) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// func (b SQLJSONMODEL) MarshalJSON() ([]byte, error) {
// 	fmt.Println("")
// 	fmt.Println("")
// 	fmt.Println("MarshalJSON", b)
// 	fmt.Println("")
// 	fmt.Println("")

// 	return json.Marshal(b)
// }
// func (b *SQLJSONMODEL) UnmarshalJSON(src []byte) error {

// 	fmt.Println("")
// 	fmt.Println("")
// 	fmt.Println("UnmarshalJSON", string(src))
// 	fmt.Println("")
// 	fmt.Println("")

// 	return json.Unmarshal(src, &b)
// }

type SQL_JSON_MODEL_ARRAY struct{}

func (b SQL_JSON_MODEL_ARRAY) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), &b)
}
func (b SQL_JSON_MODEL_ARRAY) Value() (driver.Value, error) {
	return json.Marshal(b)
}
