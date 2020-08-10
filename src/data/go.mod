module ellis.io/data

go 1.14

replace (
	ellis.io/data v0.0.0 => ./
	ellis.io/crypto/rsa v0.0.0 => ../crypto/rsa
	ellis.io/crypto/keys v0.0.0 => ../crypto/keys
)

require (
	ellis.io/crypto/rsa v0.0.0
	ellis.io/crypto/keys v0.0.0
)


