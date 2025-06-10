![Go](https://img.shields.io/badge/Go-1.20+-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

# Easypass

Easypass is a command line tool that will allow you to generate, save, and retrieve passwords for any platform, website, or anything else that you need a password for. All passwords are encrypted and securely stored in a local database so you can be assured they are safe. Passwords that are generated, updated or retrieved are automatically copied to the clipboard for easy pasting and no transcription errors.

## Why Easypass?
I kept getting crazy messages that my passwords that were saved on my browser were exposed in a data leak. I hate changing passwords because I can never remember if I updated them or not. I hated trying to think of new passwords for the many sites that I visited. I needed a tool that I could use to make my life easier so I created Easypass. With Easypass, I don't have to expose saved passwords online. I don't have to worry about my passwords getting exposed and if they are, I can easily update and manage them locally. 

## Installation

    go install github.com/rpowelson12/Easypass

This installs the program within your PATH and allows you to use the commands anywhere in your file system.

```{"db_url":"postgres://<system_name>:@localhost:5432/easypass?sslmode=disable","current_user_name":""}```

Ensure there is a file named:
```.easypassconfig.json``` 

in the root of your system with the above code pasted in. Ensure the postgres url is correct and set up with your system name and you can enter the database with that url with:

```psql "postgres://<system_name>:@localhost:5432/easypass"```

## ✨ Features
- 🔐 Secure password storage (using go's bcrypt adaptive hashing)
- 💾 Local-first, no cloud
- 📋 Clipboard integration
- 🧠 Easy to remember commands
- 🧹 Simple user and platform cleanup


## Commands

    Easypass help

Lists all available commands and a short description of what they do

    Easypass register <username>

Registers given username. Will ask for a password before completing the registration.

    Easypass login <username>

Logs in the given username. Will ask for a password after.

    Easypass users

Lists all users that are registered on your device.

    Easypass generate <platform name>

Generates a password for the given platform and copies it to clipboard

    Easypass get <platform name>

Gets the password of given platform and copies it to clipboard.

    Easypass platforms

Lists all platforms with a saved password for the currently logged in user

    Easypass delete <platform>

Deletes the platform and passwords for the current user

    Easypass deactivate <username>

Deletes the given username and deletes all their information

    Easypass update <platform name>

Updates the password for the given platform and copies it to clipboard.
