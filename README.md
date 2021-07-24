# goBotBoilerplate
Boilerplate code for writing a telegram bot using gotgbot library

[![GO](https://forthebadge.com/images/badges/made-with-go.svg)](https://golang.org/)

IGNORE THIS REPO PLEASE!!!

## Note
1. This is just a boilerplate code to make a bot with gotgbot library.
2. It's not meant to be deployed as is, i created this just so i don't
have to start from scratch everytime when i want to write a new bot.


## Development
1. By default, it only has one MessageHandler, for /start command, replies with bot's uptime.
2. Bot token can either be passed via Environmental variable, `TOKEN`
or via the first argument while running the bot.
3. To create new modules, just create a new go file in modules/ folder and then add appropriate handlers in main.go

## Build

```sh
$ git clone https://github.com/TheHamkerCat/goBotBoilerplate
$ go build .
```
## Run

```sh
$ ./goBotBoilerplate {botToken}
```

or run without any args if you've exported ENV vars
