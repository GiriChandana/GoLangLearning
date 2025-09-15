module example.com/learn

go 1.25.1

require (
	github.com/GiriChandana/GoLangLearning/cryptit v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/GiriChandana/GoLangLearning/cryptit => ../cryptit
