## Go FX Socials

A bot to watch for bad social links posted in discord, and fix them]

See: 
* [FxTwitter](https://github.com/FixTweet/FxTwitter)
* @TODO add others here

## Installation

1) Clone the repo
2) Install dependencies
3) Create a .env file with the following variables:
```
DISCORD_TOKEN=
```
4) Run the bot
```
go run main.go
```

## Discord Developer Permissions

Oauth2 Scopes
- applications.commands
- bot
  - read messages/view channels
  - send messages
  - send messages in threads

Bot Permissions
- MESSAGE CONTENT INTENT

## Usage

Host this bot somewhere, or run it locally

### Add bot user to channel

Any channel the bot has permission to view it will automatically watch for links to fix. Discord doesn't allow bots to update user messages, repost it with the fixed link.

### Slash Command

`/fix-social <url>` - The bot will fix the link you provide and repost it in the channel you're in.

## Support links

### Twitter

fx/vx - add either fx before 'twitter' to embed the video or images!

### X (gonna give it to ya)
Fix vX - If we're ever forced to use x.com, stick fixv in front of the x.

### TikTok  
vxTikTok - put vx in front of tiktok.

### Reddit
vxReddit - get rid of that fake play button by putting vx in front of reddit!

### Instagram
Insta Fix - put dd in front of Instagram. This doesn't work 100% of the time, sometimes it only works on popular videos.