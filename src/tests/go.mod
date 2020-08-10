module ellis.com/tests

go 1.14

require ellis.com/command v0.0.0

replace (
	ellis.com/command v0.0.0 => ../command
	ellis.com/data v0.0.0 => ../data
	ellis.com/crypto/rsa v0.0.0 => ../crypto/rsa
	ellis.com/crypto/keys v0.0.0 => ../crypto/keys
)