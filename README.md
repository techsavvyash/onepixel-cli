# onepixel-cli
CLI client for the onepixel backend.

Currently this is a very crude version of the cli client for [onepixel-backend](https://github.com/championswimmer/onepixel_backend), which I created as a small exercise to learn go and also to facilitate the no bullshit `1px.li` url shortener, because short shortening is a need and bit.ly has just started to suck in their free tier.
 
## Present Features
Build the app using
```bash
go build -o 1px-cli
```

1. Login using the login command
```bash
./1px-cli login
```

2. Creating short links:
> For random slugs
```bash
./1px-cli url --data "https://github.com/techsavvyash"
```

> For predefined slugs
```bash
./1px-cli url --data "https://github.com/techsavvyash" --short "git"
```

## Areas to be improved

1. Prompt formatting and beautification
2. Better Error handling
3. Code refactoring