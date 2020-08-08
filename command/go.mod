module ellis.io/command

go 1.14

require (
	ellis.io/crypto/rsa v0.0.0
	ellis.io/crypto/keys v0.0.0
)

replace (
	ellis.io/crypto/rsa v0.0.0 => ../crypto/rsa
	ellis.io/crypto/keys v0.0.0 => ../crypto/keys
)
