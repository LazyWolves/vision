package apiDoc

var ApiDoc string

func prepareDoc() {

	// The doc body
	apiDocString := `
		GET options:
			path 		:	String 	Absolute path of the resource file on the remote system
			readFrom 	:	String 	Specifies from where to read the file, can be wither of head ot tail
			limit 		:	Integer 	Specifies the number of lines to be read
			filterBy 	:	String (regex) 	A regex to filter out desired lines from the given file. Only thoe lines containing patterns matched by the given regex will be returned.
			ignore 		:	String (regex) 	A regex to exclude lines containing patterns matching the regex
			alias 		:	String 	An alias name. Must be configured beforehand
	`

	ApiDoc = apiDocString
}

func init() {
	prepareDoc()
}


