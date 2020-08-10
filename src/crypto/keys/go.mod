module ellis.io/crypto/keys

go 1.14

replace (
	ellis.io/crypto/keys v0.0.0 => ./
	ellis.io/crypto/rsa v0.0.0 => ../rsa
)

require (
	ellis.io/crypto/rsa v0.0.0
	github.com/google/uuid v1.1.1
)

