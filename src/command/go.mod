module ellis.io/command

go 1.14

replace (
	ellis.io/command v0.0.0 => ./
	ellis.io/data v0.0.0 => ../data
	ellis.io/crypto/keys v0.0.0 => ../crypto/keys
	ellis.io/crypto/rsa v0.0.0 => ../crypto/rsa
)

require (
	ellis.io/data v0.0.0
	ellis.io/crypto/keys v0.0.0
)

