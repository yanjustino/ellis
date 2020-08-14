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
  <img width="100%"  src="https://raw.githubusercontent.com/yanjustino/ellis/master/assets/lifecycle.png">
</p>

This image describes the following workflow. The person A (Alice 👩) generates two cryptography keys (PEM an JWK), take the private key and sends the public key (JWK) to Person B (Bob 👨). Now Bob, with public key, can register secret keys. After that, Bob can list and preview the encryhpted secrets. Fanilly, Bob can generate a encrypted settings file  and send it to Alice. These steps describe the ellis life cycle

# Usage
<table>
    <thead>
        <tr>
            <th>Actor</th>
            <th>Action</th>
            <th>Command</th>
            <th>Result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td width='5%' align='center'> 👩 </td>
            <td width='25%'>Creating encryption keys</td>
            <td width='35%'><code><strong>ellis keys -g [label]</strong></code> </td>
            <td width='35%'>Files 🔑[PEM] and 🔑{JWK}</td>
        </tr>
        <tr>
            <td align='center'> 👩 </td>
            <td colspan='3' align='center'>Store 🔑[PEM]</td>
        </tr>
        <tr>
            <td align='center'> 👩 </td>
            <td align='center' colspan='3'>Send 🔑{JWK} to 👨</td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td align='center' colspan='3'>Recives 🔑{JWK} from 👩</td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td>Set a secret</td>
            <td> <code><strong>ellis set -k [jwk-file] "key" "value"</strong></code> </td>
            <td>store the [🔒 secret]</td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td>List secrets ⚠️ </td>
            <td> <code><strong>ellis list -k [jwk-file]</strong></code> </td>
            <td>list of [🔒 secret]</td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td>Preview secrets ⚠️ </td>
            <td> <code><strong>ellis view -k [jwk-file]</strong></code> </td>
            <td>preview [label].settings.json 📄[JSON] </td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td>Create settings file</td>
            <td> <code><strong>ellis eject -k [jwk-file]</strong></code> </td>
            <td>creates [label].settings.json 📄[JSON] </td>
        </tr>
        <tr>
            <td align='center'> 👨 </td>
            <td align='center' colspan='3'>Send settings file 📄[JSON] to 👩</td>
        </tr>
        <tr>
            <td align='center'> 👩 </td>
            <td align='center' colspan='3'>Recives settings file 📄[JSON] from 👨 </td>
        </tr>
    </tbody>
</table>

⚠️ optional command

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
James Henry Ellis (25 September 1924 – 25 November 1997) was a British engineer and **cryptographer**. In 1970, while working at the Government Communications Headquarters (GCHQ) in Cheltenham, he conceived of the possibility of "non-secret encryption", more commonly termed **public-key cryptography**. He had a simple but clever idea: 

> Closing and opening are inverse operations. In that sense "Alice" (:girl: :key: :unlock:) could open a lock and keep the key and send the open lock to "Bob". Bob (:boy: :lock: :email:) closes his message and sends it back to Alice (:girl: :key: :unlock: :email:). No keys are exchanged. This means that she (:girl: :unlock: :unlock: :unlock: :unlock:) could publish the padlock widely and allow any person to send a message, having to keep track of only one key.

# References
* [RSA encryptation](https://tools.ietf.org/html/rfc3447)
* [JSON Web Key (JWK)](https://openid.net/specs/draft-jones-json-web-key-03.html)
* [Safe storage of app secrets](https://docs.microsoft.com/en-us/aspnet/core/security/app-secrets?view=aspnetcore-3.1&tabs=windows)

