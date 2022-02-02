package orm

// RegisterModel register models
func RegisterModel(models ...interface{}) {
	RegisterModelWithPrefix("", models...)
}

// RegisterModelWithPrefix register models with a prefix
func RegisterModelWithPrefix(prefix string, models ...interface{}) {
	if err := modelCache.register(prefix, true, models...); err != nil {
		panic(err)
	}
}

// RegisterModelWithSuffix register models with a suffix
func RegisterModelWithSuffix(suffix string, models ...interface{}) {
	if err := modelCache.register(suffix, false, models...); err != nil {
		panic(err)
	}
}

// BootStrap bootstrap models.
// make all model parsed and can not add more models
func BootStrap() {
	modelCache.bootstrap()
}
