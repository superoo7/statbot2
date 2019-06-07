# StatBot v2

StatBot is a discord bot for querying cryptocurrency pricing, operate steem blockchain related operation and many more!

# Stack

- Go
- Docker
- Nodejs
- GRPC
- AWS S3

Go would be the main connection to Discord WS, whereas nodejs will be uploading images to S3 buckets.

# Installation

There's 2 part of installation, 1 for discord bot server, and the other one is nodejs grpc server. The go repo will act as a grpc client to get data from nodejs grpc server.

## Step 0: Setup env

Move [.env.sample](.env.sample) to `.env` and add in discord token and s3 bucket token

```
cp .env.sample .env
```

## Step 1: Installation of NodeJS GRPC Server

```
cd nodejs
npm install
cd ..
```

## Step 2: Start server

Open 2 terminal and run following scripts:

```
go run main.go
(cd nodejs && node index.js)
```

# Implement into your server

Currently statbot has a [whitelist](config/whitelist.json) for limited discord channel only.

Feel free to make a PR to implement it into your server.
