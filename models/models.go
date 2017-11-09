package models

const (
	//
	// IDLength is the length of the ID parameter.
	//
	IDLength = 64

	//
	// IDRegexp is a regular expression for the ID field.
	//
	IDRegexp = "[0-9a-z]{64}"

	//
	// IDOnlyRegexp is a regular expression for only the ID value.
	//
	IDOnlyRegexp = "^[0-9a-z]{64}$"
)
