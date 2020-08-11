<p align="right">
  <img height="300" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/image.png">
</p>


# About ELLIS
Ellis is a encrypted secrets builder. It uses de concept of public-key cryptography to produces settings file with this secrets.  See how ellis do it:

## Lifecycle
<p align="center">
  <img height="600" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/lifecycle.png">
  <label>Figure 1. The ellis lifecycle</label>
</p>

The Figure 1 illustrates the following steps and their respective commands:

The Person' (Alice) uses the command  `ellis keys -g {jwk-file-name}` to generate de key-pair. Alice Take the private key and sends the public key (JWK) to Person'' (Bob). Now Bob, with public key, can register secret keys using de command `ellis set -k {jwk-file-name} [key] [value]`. After that, Bob can list ther registerd keys with the command `ellis list -k {jwk-file-name}`. The result will be like this:

```
[0] key: keyA - value: DdQ5brEeK8lYyT0g72OUnrkVlbDUu0UYZu0W67U9EOvxGkjXVVWTQ3Mm6iGbJBroG+8cOWM7y2ILbugR3uL5um5aayMUpKCDEtWG32IycfgVeVWV0kgv0tLeiaIQ8bY5IzUQtRUUizaPK8ereHSWR1tsfW6cPrXo+vAUwgGxJQaGnip3JZwGzZZXO8Sx9Mb/3PLcODEePBNpFzyD3ZX7kgE5UknbctcmFp25Uj+BCWZhghis1noOMG6y8dfMPjf0H6KTO81Odsch54xKhML9fy+8Fw40IbcTVCORzJfx1JvjTddHWHKo+606JDbaYvOqGUTi8oOTSTUs0imV892M7Q==
[1] key: keyB - value: eB2OH+r6B6K1WW79vY+2kosxewlc2cyeDNGhT87pyH1AE5rHdqIcm3SsTbYXgyyxTotGOh1+VmwcRonJ3K+3jFZ7oA4ELpU+b4mASxy3L4wvmiaQap23nPkYbk3BMedlh4vE6u0u2eXH3qnHXUKBN7HiGKcEyqhHHy87eq9a/3RnajFKvkUsLmcRB4zP1CTpAfQWvPxs61J2EzqIJTlrqk4qWY+0A97yVCTjrhFWuU/3zD4Ip9wMSaCeMyhcflv7wG+hv1RYboH9i0tXUInDfcuX67CtHcl3YHjGDT5o8Zq6ilWEPfemV4deKxsOc0ABpSflAvel+nCxEJM5ra27IA==
```
Alson, Bob can view a preview result of final settings file (Secrets Holder) using the command `ellis view -k {jwk-file-name}` :

```
{
 "Items": [
  {
   "key": "keyA",
   "value": "DdQ5brEeK8lYyT0g72OUnrkVlbDUu0UYZu0W67U9EOvxGkjXVVWTQ3Mm6iGbJBroG+8cOWM7y2ILbugR3uL5um5aayMUpKCDEtWG32IycfgVeVWV0kgv0tLeiaIQ8bY5IzUQtRUUizaPK8ereHSWR1tsfW6cPrXo+vAUwgGxJQaGnip3JZwGzZZXO8Sx9Mb/3PLcODEePBNpFzyD3ZX7kgE5UknbctcmFp25Uj+BCWZhghis1noOMG6y8dfMPjf0H6KTO81Odsch54xKhML9fy+8Fw40IbcTVCORzJfx1JvjTddHWHKo+606JDbaYvOqGUTi8oOTSTUs0imV892M7Q=="
  },
  {
   "key": "keyB",
   "value": "eB2OH+r6B6K1WW79vY+2kosxewlc2cyeDNGhT87pyH1AE5rHdqIcm3SsTbYXgyyxTotGOh1+VmwcRonJ3K+3jFZ7oA4ELpU+b4mASxy3L4wvmiaQap23nPkYbk3BMedlh4vE6u0u2eXH3qnHXUKBN7HiGKcEyqhHHy87eq9a/3RnajFKvkUsLmcRB4zP1CTpAfQWvPxs61J2EzqIJTlrqk4qWY+0A97yVCTjrhFWuU/3zD4Ip9wMSaCeMyhcflv7wG+hv1RYboH9i0tXUInDfcuX67CtHcl3YHjGDT5o8Zq6ilWEPfemV4deKxsOc0ABpSflAvel+nCxEJM5ra27IA=="
  }
 ]
}
```
Then, Bob can generate this Json as a File using the command `ellis eject -k {jwk-file-name}` and send it to Alice, who can decrypt this file with his key. 
These steps describe the ellis life cycle


## Usage
```shell
 Usage: ellis <command> [options] [path-to-jwk]
 Usage: ellis [path-to-jwk]
 
 * command:
   * help  Display help 
   * keys  Generate a pair of RSA Keys (The public key is generated in JWK format) 
   * list  List all secrets for JWK file 
   * view  Preview the settings file 
   * set   Store a key-value pair 
   * eject Generate a settings file 
 
 * command [options]:
   * help  [-h] 
   * keys  [-g] [jwk-id] 
   * list  [-k] [path-to-jwk] 
   * view  [-k] [path-to-jwk] 
   * set   [-k] [path-to-jwk] [key] [value]
   * eject [-k] [path-to-jwk] 
 
* path-to-jwk:
 	* The path to an application JWK file to execute.
  ```
  
# Troble storing secrets
Production secrets shouldn't be used for development or test. In addition, secrets should not be implanted with the application. Now consider that your application needs to run in different environments and that for each one it needs a different secret. In this context, your application needs to manage different keys or it will carry several messages with secret managers to establish a communication. Was there a simple way?

James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**.

He had a simple but clever idea: Closing and opening are inverse operations. In that sense "Alice" (:girl: :key: :unlock:) could open a lock and keep the key and send the open lock to "Bob". Bob (:boy: :lock: :email:) closes his message and sends it back to Alice (:girl: :key: :unlock: :email:). No keys are exchanged. This means that she (:girl: :unlock: :unlock: :unlock: :unlock:) could publish the padlock widely and allow any person to send a message, having to keep track of only one key.
  

# Who was James H. Ellis 
James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**.


# References
* [RSA encryptation](https://www.khanacademy.org/computing/computer-science/cryptography/modern-crypt/v/intro-to-rsa-encryption)
