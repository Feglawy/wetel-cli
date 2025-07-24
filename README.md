# wetel-cli

**wetel-cli** is a lightweight command-line tool for accessing your [We Telecom](https://te.eg/) account information directly from the terminal.

Easily check your service number, balance, data usage, and plan details — all without needing to open a browser.

---

## ✨ Features

- ✅ Login with your We Telecom service number and password
- 📊 View current balance and plan usage
- 💡 Check remaining and used data
- 📦 See active offers and detailed plan info

---

## 🚀 Installation

### Go (recommended)

```bash
go install github.com/Feglawy/wetel-cli/cmd/wetel-cli@latest
```

Make sure your Go binary path is in your PATH. Usually it's:
- Linux/macOS: `~/go/bin`
- Windows: `%USERPROFILE%\go\bin`

##  Usage
once you installed it
```bash
wetel-cli
```
Available Flags
| Flag    | Description                                          | Example                 |
| ------- | ---------------------------------------------------- | ----------------------- |
| `-num`  | Service number for login (e.g., your phone number)   | `-num=0238900000`       |
| `-pass` | Password for login                                   | `-pass=mypass123`       |
| `-r`    | Remember login credentials for future use (optional) | `-r`                    |

example:
```bash
wetel-cli -num=0238900000 -pass=mypass123 -r
```
If no flags are provided, the CLI will prompt you to enter them interactively, with the password hidden during input for security.

## Configuration
If `-r` is enabled, credentials are stored in your system in a file beside the app called `loginInfo.json`
you can edit it or delete it so you can login for other accounts

## 🧰 Requirements

- Go 1.18+
- An active [We Telecom account](https://my.te.eg/echannel/#/login)
- Internet access

## 🧾 License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.

## 🙋 FAQ
❓ Is this an official We Telecom tool?
- No — this is an independent, open-source CLI made for convenience.
- 
## 🤝 Contributing

Pull requests and suggestions are welcome!
1. Fork the repo
2. Create a new branch
3. Commit your changes
4. Open a PR