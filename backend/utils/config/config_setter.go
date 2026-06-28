package config

// SetConfigString set the config or return the value in field directly
func SetConfigString(
	key string,
	defaultValue string,
	field **string,
) string {
	if field == nil {
		panic("You cannot give a nil pointer to field in SetConfigString!")
	}
	if *field == nil {
		*field = GetEnvString(
			key,
			defaultValue,
		)
	}
	return **field
}
