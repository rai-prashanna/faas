package config

func Load() map[string]interface{} {
	config := map[string]interface{}{
		"Website": "Dig Service website",
		"host": "localhost",
		"port": "9090",
	}

	return config
}
