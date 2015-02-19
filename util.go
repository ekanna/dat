package dat

import "bytes"

// NameMapping is the routine to use when mapping column names to struct properties
var NameMapping = camelCaseToSnakeCase

func camelCaseToSnakeCase(name string) string {
	var buf bytes.Buffer

	firstTime := true
	for _, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime {
				firstTime = false
			} else {
				buf.WriteRune('_')
			}
			chr -= ('A' - 'a')
		}
		buf.WriteRune(chr)
	}

	return buf.String()
}

func camelCaseToSnakeCaseID(name string) string {
	// handle the common ID idiom
	if name == "ID" {
		return "id"
	}
	var buf bytes.Buffer

	lenName := len(name)
	writeID := false
	if lenName > 2 {
		writeID = name[lenName-2:lenName-1] == "I" && name[lenName-1:lenName] == "D"
	}

	firstTime := true
	for i, chr := range name {
		if writeID && i == lenName-2 {
			buf.WriteString("_id")
			break
		}
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime {
				firstTime = false
			} else {
				buf.WriteRune('_')
			}
			chr -= ('A' - 'a')
		}
		buf.WriteRune(chr)
	}

	return buf.String()
}