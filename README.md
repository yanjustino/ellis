<p align="right">
  <img height="300" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/image.png">
</p>

Production secrets shouldn't be used for development or test. In addition, secrets should not be implanted with the application. Now consider that your application needs to run in different environments and that for each one it needs a different secret. In this context, your application needs to manage different keys or it will carry several messages with secret managers to establish a communication. Was there a simple way?

# About ELLIS
Ellis is a simple secret encrypter. It uses de concept of public-key cryptography to produces encrypted data json file to using in application settings file in distincts environment. See how ellis do it:

## Lifecycle
<p align="center">
  <label><strong>Figure 1.</strong> The ellis lifecycle</label>
  <img height="600" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/lifecycle.png">
</p>

The Figure 1 illustrates the following steps and their respective commands:

The Person' (Alice :girl:) uses the command  **`ellis keys -g {jwk-file-name}`** to generate de key-pair. Alice Take the private key and sends the public key (JWK) to Person'' (Bob :boy:). Now Bob, with public key, can register secret keys using de command **`ellis set -k {jwk-file-name} [key] [value]`**. After that, Bob can list ther registerd keys with the command **`ellis list -k {jwk-file-name}`**. The result will be like this:

```
[0] key: keyA - value: DdQ5brEeK8lYyT0g72OUnrkVlbDUu0UYZu0W67U9EOvxGkjXVVWTQ3...
[1] key: keyB - value: eB2OH+r6B6K1WW79vY+2kosxewlc2cyeDNGhT87pyH1AE5rHdqIcm3...
```
Alson, Bob can view a preview result of final settings file (Secrets Holder) using the command `ellis view -k {jwk-file-name}` :

```
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
Then, Bob can generate this Json as a File using the command **`ellis eject -k {jwk-file-name}`** and send it to Alice, who can store and decrypt this file with his key. 
These steps describe the ellis life cycle


# Who was James H. Ellis 
James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**.

He had a simple but clever idea: Closing and opening are inverse operations. In that sense "Alice" (:girl: :key: :unlock:) could open a lock and keep the key and send the open lock to "Bob". Bob (:boy: :lock: :email:) closes his message and sends it back to Alice (:girl: :key: :unlock: :email:). No keys are exchanged. This means that she (:girl: :unlock: :unlock: :unlock: :unlock:) could publish the padlock widely and allow any person to send a message, having to keep track of only one key.

# References
* [RSA encryptation](https://www.khanacademy.org/computing/computer-science/cryptography/modern-crypt/v/intro-to-rsa-encryption)
