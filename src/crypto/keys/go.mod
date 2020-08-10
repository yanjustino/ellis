module ellis.com/crypto/keys

go 1.14

require (
	ellis.com/crypto/rsa v0.0.0
	github.com/google/uuid v1.1.1
)

replace (
	ellis.com/crypto/keys v0.0.0 => ./
	ellis.com/crypto/rsa v0.0.0 => ../rsa
)

