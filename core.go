package articles

// Repository inteface
type Repository interface {
	// Start starts the repository
	Start()

	// Close closes any resources used by the repository
	Stop()
}

// Service inteface
type Service interface {
	// Start starts the repository
	Start()

	// Close closes any resources used by the repository
	Stop()
}
