# Configuring Slack

## Create Slack App
- Access https://api.slack.com/apps?new_app=1 in your workspace
- Click on the button: __Create an App__
- A model will be show to you, so fit with your __App Name__ and select your __Slack Workspace__
- Click on the button: __Create App__
- In __Features__ click on the __OAuth & Permissions__
- In __Scopes__/__Bot Token Scopes__ click on the __Add an OAuth Scope__
- Then select the follows scopes:


### Bot Token Scopes
```
 - app_mentions:read
 - channels:join
 - channels:manage
 - channels:read
 - channels:history
 - chat:write.public
 - chat:write
 - commands
 - groups:read
 - files:read
 - links:read
 - mpim:write
 - pins:read
 - pins:write
 - usergroups:read
 - users:read
 - users:read.email
```

### User Token Scopes
```
- channels:read
- channels:write
- chat:write
- mpim:write
- pins:read
- users:read
- users:read.email
```

4. Install app to workspace
- In __Features__ click on the __OAuth & Permissions__
- And click on the __Install App to Workspace__
- Select a channel to Slack post to as an app and click on the __Allow__
- Copy the __Bot User OAuth Access Token__
- Paste de code into variables `HELLPER_OAUTH_TOKEN`
- Restart your application


5. Events
- Now, in __Features__ click on the __Event Subscriptions__
- And, in __Enable Events__ turn on it
- In __Request URL__, fit the filed with your application's URL public. It will look something like this: `https://yourhost.publicaddress.com/events`
- Under it you should receive a request look like this:
```
Our Request:
POST
"body": {
	 "type": "url_verification",
	 "token": "XXXXXXXX",
	 "challenge": "YYYYYYY"
}
Your Response:
"code": 200
"error": "challenge_failed"
"body": {

}
```
_* in your console application will receive the same info too._

- Copy the `token` and past into `hellper_verification_token`.
- Restar your application and try again


Subscribe to bot events

- In the same page open the __Subscribe to bot events__, click on the __Add Bot User Event__ and add `app_mention` option.
- After that, open the __Subscribe to events on behalf of users__, click on the __Add Workspace Event__ and add `channel_created` option.
- Now, go in __Features__, click on the __Slash Commands__ and click on the __Create New Command__ to add the follows commands:


| Command  | Request URL | Short Description |
| - | - | - |
|`/hellper_incident`|https://44744f1c.ngrok.io/open|_Start Incident_|
|`/hellper_status`|https://44744f1c.ngrok.io/status|_Show all pinned messages_|
|`/hellper_close`|https://44744f1c.ngrok.io/close|_Close Incident_|
|`/hellper_resolve`|https://44744f1c.ngrok.io/resolve|_Resolve Incident_|
|`/hellper_cancel`|https://44744f1c.ngrok.io/cancel|_Cancel Incident_|
|`/hellper_update_dates`|https://44744f1c.ngrok.io/dates|_Update the dates for an incident_|

- At last, click on the __Save Changes__

- Now, in __Features__/__Interactivity & Shortcuts__ turn on the option __Interactivity__ and configure your address URL `http://44744f1c.ngrok.io/interactive`