<p align="right">
  <img height="300" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/image.png">
</p>

# Troble storing secrets
Production secrets shouldn't be used for development or test. In addition, secrets should not be implanted with the application. Now consider that your application needs to run in different environments and that for each one it needs a different secret. In this context, your application needs to manage different keys or it will carry several messages with secret managers to establish a communication. Was there a simple way?

James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**.

He had a simple but clever idea: Closing and opening are inverse operations. In that sense "Alice" (:girl: :key: :unlock:) could open a lock and keep the key and send the open lock to "Bob". Bob (:boy: :lock: :email:) closes his message and sends it back to Alice (:girl: :key: :unlock: :email:). No keys are exchanged. This means that she (:girl: :unlock: :unlock: :unlock: :unlock:) could publish the padlock widely and allow any person to send a message, having to keep track of only one key.

# About ELLIS
Ellis is a encrypted secrets builder. It uses de concept of **public-key cryptography** to produces settings file with this encrypted secrets. Like the Alice and Bob´s strory, **ellis** allows any application recive encrypted settins, having to keep track of only one key. See how ellis do it:

## Lifecycle
<p align="center">
  <img height="650" src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/lifecycle.png">
</p>

# Who was James H. Ellis 
James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**.


# References
* [RSA encryptation](https://www.khanacademy.org/computing/computer-science/cryptography/modern-crypt/v/intro-to-rsa-encryption)
