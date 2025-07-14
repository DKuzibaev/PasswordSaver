# PasswordSaver

A simple command-line password manager written in Go. Allows users to create, find, and delete account credentials securely.

## Features
- Create new accounts with login, password, and URL.
- Find accounts by URL.
- Delete accounts (functionality not implemented).
- Colorful console interface using the `fatih/color` package.

## Prerequisites
- Go 1.16 or higher
- Required package: `github.com/fatih/color`

## Installation
1. Clone the repository:
   ```git clone <repository-url>```
2. Install dependencies:
```go get github.com/fatih/color```
```go run main.go```

## Usage

## Run the program and choose an option from the menu:
    Create Account: Enter login, password, and URL to store a new account.
    Find Account: Search for an account by URL.
    Delete Account: (Not implemented) Delete an existing account.
    Exit: Close the program.

## Notes
    The account package is assumed to handle account creation, storage, and retrieval logic.
    The deleteAccount function is currently a placeholder and not implemented.
    Input validation ensures proper URL format and handles input errors.

## License
  MIT License
