# Sealenv: Environment Variables Encryption Tool

## Overview

Sealenv is a command-line tool written in Go that allows you to securely encrypt and decrypt your environment variables. It helps you execute commands with sensitive data without exposing it in plain text. Sealenv is handy for managing environment variables in local and remote environments to pass credentials and any other sensivity data to CLI programs like **psql**, **SnowSQL**, **trino**, etc where credential can be set by environment variables

## Features

- **Run Commands Securely**: Execute commands using encrypted environment variables securely. You're asked for your master password to be used to decrypted the environment file and run the command. The decrypted variables are never store in the file system
- **Encrypt Environment Variables**: Safeguard your environment variables by encrypting your **.env** files.
- **Decrypt Environment Variables**: Show the decrypted values of your **.env** file.
- **Portable**: Written in Go, Sealenv compiles to a single binary for easy distribution to every OS/Arch supported by Go.

## Prerequisites

- Go 1.18+ (for building from source)

## Installation

### From Source

1. Clone the repository:

   ```bash
   git clone https://github.com/ulisesojeda/sealenv
   ```

2. Build the binary:

   ```bash
   go build main.go -o sealenv
   ```

3. Move the binary to your `$PATH`:

   ```bash
   mv sealenv /usr/local/bin/ # Linux
   mv sealenv ~/bin/ # MacOS
   ```

   For **Windows**, copy the binary to any folder contained in your system **PATH**

### Using a Precompiled Binary

1. Download the latest release from the [releases page](https://github.com/ulisesojeda/sealenv/releases).

2. Move the binary to your `$PATH`:

   ```bash
   mv sealenv_linux_amd64 /usr/local/bin/sealenv # Linux
   mv sealenv_darwin_amd64 ~/bin/sealenv # MacOS
   ```

   For **Windows**, copy the binary **windows_amd64.exe** to any folder contained in your system **PATH**

## Usage

### Encrypting Environment Variables

1. Create a file (e.g: **creds.plain**) with your variables in this format:

```
   PGUSER=myuser
   PGPASSWORD=mysecretpassword
```

2. Encrypt your environment variables using the following command:

```bash
sealenv encrypt --env creds.plain --out .env
```

This command will create a file **.env** with all variables encrypted with your master password

3. Define your master password. **IMPORTANT**: this password will be used to encrypt/decrypt your variables and will asked every time run a command with **sealenv**. If you forget the master password the credentials couldn't recovered.

### Decrypting Environment Variables

```bash
sealenv decrypt --env .env
```

This command will print the decrypted values of your environment variables

### Run a program with encrypted environment variables

```bash
sealenv run 'psql -h localhost -d postgres'  --env .env
```

### How does it works ?

The environment variables plain text are encrypted using AES-GCM simmetric-key algorithm.
The encrypted values are store in the file system and decrypted in-memory upon program execution.

### License

This project is licensed under the [MIT License](https://opensource.org/license/mit).
