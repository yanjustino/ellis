module ellis.com/data

go 1.14

replace (
	ellis.com/data v0.0.0 => ./
	ellis.com/crypto/rsa v0.0.0 => ../crypto/rsa
	ellis.com/crypto/keys v0.0.0 => ../crypto/keys
)

require (
	ellis.com/crypto/rsa v0.0.0
	ellis.com/crypto/keys v0.0.0
)


