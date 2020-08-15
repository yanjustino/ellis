<p align="right">
  <img height="300" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/image.png">
</p>

Ellis is a simple secret encrypter. It uses de concept of public-key cryptography to produces encrypted data json (JWK) file to using in application settings file. See how ellis do it:

![Travis (.org)](https://img.shields.io/travis/yanjustino/ellis)
[![analysis](https://img.shields.io/codeclimate/maintainability/yanjustino/ellis)](https://codeclimate.com/github/yanjustino/ellis)
![GitHub repo size](https://img.shields.io/github/repo-size/yanjustino/ellis)
![GitHub last commit](https://img.shields.io/github/last-commit/yanjustino/ellis)
![GitHub (Pre-)Release Date](https://img.shields.io/github/release-date-pre/yanjustino/ellis)

## Lifecycle
<p align="center">
  <label><strong>Figure 1.</strong> The ellis lifecycle</label>
  
  <img width="80%"  src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/diagram.png">
</p>

This image describes the following workflow. The person A (Alice 👩) generates two cryptography keys (PEM an JWK), take the private key and sends the public key (JWK) to Person B (Bob 👨). Now Bob, with public key, can register secret keys. After that, Bob can list and preview the encryhpted secrets. Fanilly, Bob can generate a encrypted settings file  and send it to Alice. These steps describe the ellis life cycle.

# Getting Started

### Generate Key
```
ellis keys -g [label]
```
This command will create two files: 🔑[PEM] and 🔑{JWK}

### Encrypt Secrets
```
ellis set -k [jwk-file] "key" "value"
```
This command stores the secret (🔑{JWK}[🔒 secret])

### List or View Encrypted Secrets
```
ellis list -k [jwk-file]

Output
[0] key: keyA - value: DdQ5brEeK8lYyT0g72OUnrkVlbDUu0UYZu0W67U9EOvxGkjXVVWTQ3...
[1] key: keyB - value: eB2OH+r6B6K1WW79vY+2kosxewlc2cyeDNGhT87pyH1AE5rHdqIcm3...
```

```
ellis view -k [jwk-file]

Output
{
   "Items": [
      {
         "key": "keyA",
         "value": "DdQ5brEeK8lYyT0g72OUnrkVlbDUu0UYZu0W67U9EOvxGkjXVVWTQ3Mm6iGbJB..."
      },
      {
         "key": "keyB",
         "value": "eB2OH+r6B6K1WW79vY+2kosxewlc2cyeDNGhT87pyH1AE5rHdqIcm3SsTbYXgy..."
      }
   ]
}
```
This command shown a list of encrypted secrets (🔑{JWK}[🔒 secret])

###  Create Encrypted Setting
```
ellis eject -k [jwk-file]
```
This command creates [label].settings.json file 📄[JSON]

### Download ELLIS

You can download a realiase version [clicking here](https://github.com/yanjustino/ellis/releases). Another way is use [Make](https://www.gnu.org/software/make/) to build the ellis program:

```
$ git clone https://github.com/yanjustino/ellis.git
$ cd ellis
$ make install
```

This commands will generate ellis program in `./bin` directory. If you are a windows user, the programa will be `./bin/ellis.exe`


# Guide to contributing to a GitHub project
This is a guide to contributing to this open source project that uses GitHub. It’s mostly based on how many open sorce projects operate. That’s all there is to it. The fundamentals are:

* Fork the project & clone locally.
* Create an upstream remote and sync your local copy before you branch.
* Branch for each separate piece of work.
* Do the work, write good commit messages, and read the CONTRIBUTING file if there is one.
* Push to your origin repository.
* Create a new PR in GitHub.
* Respond to any code review feedback.

If you want to contribute to an open source project, the best one to pick is one that you are using yourself. The maintainers will appreciate it!

# Who was James H. Ellis 
James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**. 

# References
* [RSA encryptation](https://tools.ietf.org/html/rfc3447)
* [JSON Web Key (JWK)](https://openid.net/specs/draft-jones-json-web-key-03.html)
* [Safe storage of app secrets](https://docs.microsoft.com/en-us/aspnet/core/security/app-secrets?view=aspnetcore-3.1&tabs=windows)

