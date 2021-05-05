// Code generated by entc, DO NOT EDIT.

package devicerequest

const (
	// Label holds the string label denoting the devicerequest type in the database.
	Label = "device_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserCode holds the string denoting the user_code field in the database.
	FieldUserCode = "user_code"
	// FieldDeviceCode holds the string denoting the device_code field in the database.
	FieldDeviceCode = "device_code"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldClientSecret holds the string denoting the client_secret field in the database.
	FieldClientSecret = "client_secret"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// Table holds the table name of the devicerequest in the database.
	Table = "device_requests"
)

// Columns holds all SQL columns for devicerequest fields.
var Columns = []string{
	FieldID,
	FieldUserCode,
	FieldDeviceCode,
	FieldClientID,
	FieldClientSecret,
	FieldScopes,
	FieldExpiry,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserCodeValidator is a validator for the "user_code" field. It is called by the builders before save.
	UserCodeValidator func(string) error
	// DeviceCodeValidator is a validator for the "device_code" field. It is called by the builders before save.
	DeviceCodeValidator func(string) error
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// ClientSecretValidator is a validator for the "client_secret" field. It is called by the builders before save.
	ClientSecretValidator func(string) error
)
