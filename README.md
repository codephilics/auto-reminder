# Auto Reminder

Auto Reminder tool helps the user by send a notification using AWS SES when their API Keys or password are about to expire. You also can disable their account or deactivate their key if they're older than the age of the AWS IAM Account Password Policy.

# Development: Getting Started

# Requirement
* go 1.13+
* AWS 

# Prepare Workspace

Clone Project

```bash
$ git clone https://github.com/codephilics/atuo-reminder.git
$ cd atuo-reminder
```

Install the dependencies
```bash
$ go get ./...
```

To run the project

```bash
$ go run ./cmd/main.go
```

## Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.

## License
The **accounting-app** is an open-source software licensed under the [MIT License](LICENSE).
