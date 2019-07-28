package config

func Load() map[string]interface{} {
	config := map[string]interface{}{
		"Website": "Factorial website",
		"host": "localhost",
		"port": "7070",
	}

	return config
}
