# Easypass

Easypass is a command line tool that will allow you to generate, save, and retrieve passwords for any platform, website, or anything else that you need a password for. All passwords that are securely stored in the database so you can be assured they are safe. Passwords that are generated, updated or retrieved are automatically copied to the clipboard for easy pasting and no transcription errors.

## Why Easypass?
Out of so many password managers, why should you care about Easypass? Easypass is a command line tool that takes your passwords off of the internet so there is less chance of them getting exposed in a data breach. Not only are your passwords off the internet, they are further encrypted when they are stored on your local machine. Easypass quickly generates a long password mixed with special characters, uppercase and lowercase letters that satisfy the requirements for most websites! 

## Installation

    go install github.com/rpowelson12/Easypass

This installs the program within your PATH and allows you to use the commands anywhere in your file system.

```{"db_url":"postgres://<system_name>:@localhost:5432/easypass?sslmode=disable","current_user_name":""}```
Ensure there is a file named ```.easypassconfig.json``` in the root of your system with the above code pasted in.

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
