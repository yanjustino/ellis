module ellis.io

go 1.14

require (
	ellis.io/command v0.0.0
)

replace (
	ellis.io v0.0.0 => ./
	ellis.io/data v0.0.0 => ./data
	ellis.io/crypto/rsa v0.0.0 => ./crypto/rsa
	ellis.io/crypto/keys v0.0.0 => ./crypto/keys
	ellis.io/command v0.0.0 => ./command
)