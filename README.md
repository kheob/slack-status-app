# slack-status-app
Because it takes too many clicks to set your status to 🥪

## Installation
1. Add `SLACK_STATUS_APP_TOKEN` as an environment variable. To get this token you'll need to either create a a new Slack app with the `users.profile:write` scope or use an existing app that has this scope set.

2. Run `make build` to build the binary (requires Go `1.22`)

3. (Optional) Add to path (Mac users: `mv bin/slack-status /usr/local/bin`)

## Usage
```
  -emoji string
        status emoji to set
  -expires int
        status expiration in minutes (default 60)
  -status string
        status text to set
```

e.g. `./path/to/slack-status -status Hello -emoji :smile: -expires 30`
