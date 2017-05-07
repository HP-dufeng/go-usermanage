package common

// StartUp the necessary initialization logic before the HTTP server start
func StartUp() {
	// Initialize AppConfig variable
	initConfig()

	// Initialize private/public keys for JWT authentication
	initKeys()

	// Start a MongoDB session
	createDbSession()

	// Add indexes into MongoDB
	addIndexes()
}
