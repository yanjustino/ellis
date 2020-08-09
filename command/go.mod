module ellis.io/command

go 1.14

require (
	ellis.io/data v0.0.0
	ellis.io/crypto/keys v0.0.0
)

replace (
	ellis.io/utils v0.0.0 => ../utils
	ellis.io/data v0.0.0 => ../data
	ellis.io/crypto/keys v0.0.0 => ../crypto/keys
	ellis.io/crypto/rsa v0.0.0 => ../crypto/rsa
)