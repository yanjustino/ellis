module ellis.com

go 1.14

require (
	ellis.com/command v0.0.0
	ellis.com/tests v0.0.0
)

replace (
	ellis.com/command v0.0.0 => ./command
	ellis.com/crypto/keys v0.0.0 => ./crypto/keys
	ellis.com/crypto/rsa v0.0.0 => ./crypto/rsa
	ellis.com/data v0.0.0 => ./data
	ellis.com/tests v0.0.0 => ./tests
)

