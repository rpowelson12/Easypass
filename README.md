![Go](https://img.shields.io/badge/Go-1.20+-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

# Easypass

Easypass is a command-line tool that lets you generate, save, and retrieve passwords for any platform, website, or anything else that you need a password for. All passwords are encrypted and securely stored in a local database so you can be assured they are safe. Passwords that are generated, updated or retrieved are automatically copied to the clipboard for easy pasting and no transcription errors.

## 💭 Why Easypass?
After receiving multiple alerts about my browser-saved passwords being exposed in data leaks, I realized I needed a safer solution. I hated managing password updates and coming up with new ones for every site. I needed a tool that I could use to make my life easier. Easypass was born! With Easypass, I don't have to expose saved passwords online. I don't have to worry about my passwords getting exposed and if they are, I can easily update and manage them locally. 

## 📚 Table of Contents

- [Why Easypass?](#-why-easypass)
- [Install Easypass](#-install-easypass)
- [Quick Start](#-quick-start)
- [Features](#-features)
- [Demo](#-demo-terminal-example)
- [Commands](#-commands)
- [Built With](#-built-with)
- [Contributing](#-contributing)
- [License](#-license)

### 📥 Install Easypass

```bash
curl -fsSL https://raw.githubusercontent.com/rpowelson12/Easypass/main/scripts/install.sh | bash

```
Create a file named `.easypassconfig.json` in your home directory with the following content:

```json
{"db_url":"postgres://<your_username>:@localhost:5432/easypass?sslmode=disable","current_user_name":""}
```

Make sure to replace <your_username> with your system's PostgreSQL username. You can verify your connection with:

```psql "postgres://<your_username>:@localhost:5432/easypass"```


## ✨ Features
- 🔐 Secure password storage (using go's bcrypt adaptive hashing)
- 💾 Local-first, no cloud
- 📋 Clipboard integration
- 🧠 Easy to remember commands
- 🧹 Simple user and platform cleanup

## 📺 Demo (Terminal Example)
```
$ easypass register alice
Enter password:
Registration successful!

$ easypass generate github
Generated password copied to clipboard.

$ easypass get github
Password copied to clipboard.
```

## 💻 Commands

    easypass help

Lists all available commands and a short description of what they do

    easypass register <username>

Registers given username. Will ask for a password before completing the registration.

    easypass login <username>

Logs in the given username. Will ask for a password after.

    easypass users

Lists all users that are registered on your device.

    easypass generate <platform name>

Generates a password for the given platform and copies it to clipboard

    easypass get <platform name>

Gets the password of given platform and copies it to clipboard.

    easypass platforms

Lists all platforms with a saved password for the currently logged in user

    easypass delete <platform>

Deletes the platform and passwords for the current user

    easypass deactivate <username>

Deletes the given username and deletes all their information

    easypass new <platform name>

Updates the password for the given platform and copies it to clipboard.

    easypass update
    
Upgrades to the latest version of Easypass


## 🛠 Built With
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [pq](https://pkg.go.dev/github.com/lib/pq) (Postgres driver)

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/rpowelson12/Easypass@latest
cd Easypass
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## 📄 License

This project is licensed under the MIT License.

