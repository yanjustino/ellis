module ellis.com/command

go 1.14

require (
	ellis.com/data v0.0.0
	ellis.com/crypto/keys v0.0.0
)

replace (
	ellis.com/data v0.0.0 => ../data
	ellis.com/crypto/keys v0.0.0 => ../crypto/keys
	ellis.com/crypto/rsa v0.0.0 => ../crypto/rsa
)