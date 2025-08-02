# wetel-cli

    ██╗    ██╗███████╗████████╗███████╗██╗               ██████╗██╗     ██╗
    ██║    ██║██╔════╝╚══██╔══╝██╔════╝██║              ██╔════╝██║     ██║
    ██║ █╗ ██║█████╗     ██║   █████╗  ██║      █████╗  ██║     ██║     ██║
    ██║███╗██║██╔══╝     ██║   ██╔══╝  ██║      ╚════╝  ██║     ██║     ██║
    ╚███╔███╔╝███████╗   ██║   ███████╗███████╗         ╚██████╗███████╗██║
     ╚══╝╚══╝ ╚══════╝   ╚═╝   ╚══════╝╚══════╝          ╚═════╝╚══════╝╚═╝

**wetel-cli** is a lightweight command-line tool for accessing your [We Telecom](https://te.eg/) Egypt account information directly from the terminal.

Easily check your service number, balance, data usage, and plan details — all without needing to open a browser.

--- 
## 🎯 Project Goal

This project aims to provide a practical solution for users subscribed to WE Telecom who face significant slowdowns after exhausting their data quota. When the main internet quota is consumed, the connection speed is reduced to 256 kbps—insufficient for accessing the WE website or renewing the subscription online.

The objective of this project is to:

- Enable users to manage their WE account efficiently, even under limited bandwidth conditions.

- Allow users to view their current balance and quota usage.

- Provide a command-line interface (CLI) tool to renew their plan or subscribe to available add-ons—provided sufficient balance is available in the account.

**Important Note:**
    This tool does not bypass WE's systems or allow users to subscribe without having the required balance. It is strictly a convenience tool built to improve user accessibility and control.

---

## ✨ Features

- ✅ Login with your We Telecom service number and password
- 📊 View current balance and plan usage
- 💡 Check remaining and used data
- 📦 See active offers and detailed plan info
- ✨ Renew your main plan
- ✨ Subscribe to an Addon

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
- An active [We Telecom Egypt account](https://my.te.eg/echannel/#/login)
- Internet access

## 🧾 License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.

## 🙋 FAQ
❓ Is this an official WE Telecom tool?
- No — this is an independent, open-source CLI developed for convenience. It is not affiliated with, endorsed by, or supported by WE Telecom.

❓ Does this tool allow me to subscribe to a plan or add-on without having balance?
- No — you must have sufficient balance in your WE account to subscribe. This tool does not bypass WE’s systems or perform any unauthorized actions.

❓ Can I use this tool when my connection speed is reduced to 256 kbps?
- Yes — this tool is specifically designed to work under low-bandwidth conditions, making it easier to manage your account even after your main quota is depleted.

❓ What can I do with this tool?
- Check your current balance
- View remaining quota
- Subscribe to available add-ons
- Renew your current plan
- Perform basic account management via the command line

## 🤝 Contributing
Pull requests and suggestions are welcome!
1. Fork the repo
2. Create a new branch
3. Commit your changes
4. Open a PR