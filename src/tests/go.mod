module ellis.io/tests

go 1.14

require ellis.io/command v0.0.0

replace (
	ellis.io/tests v0.0.0 => ./
	ellis.io/command v0.0.0 => ../command
	ellis.io/data v0.0.0 => ../data
	ellis.io/crypto/rsa v0.0.0 => ../crypto/rsa
	ellis.io/crypto/keys v0.0.0 => ../crypto/keys
)
