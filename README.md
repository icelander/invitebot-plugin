# Mattermost Invitebot Plugin

## What is it?

This plugin automates the process of inviting guests to a Mattermost server. This link can be shared publicly, and users just have to enter their email addresses to get an invitation to join specific channels on a Mattermost server. This is great for onboarding new community members or engaging with conference attendees.


## Configuration Settings

| Setting         | Type | Default Value                                                            |
|:----------------|:-----|:-------------------------------------------------------------------------|
| LandingHTML     | Text | Stored in `assets/landing_default.html`                                  |
| UserID          | Text | ID of the user to send the invites from, must have channel admin rights  |
| ThankYouHtml    | Text | stored in `assets/thankyou_default.html`                                 |
| ErrorHTML       | Text | stored in `assets/error_default.html`                                    |
| MarketoCampaign | Text | The campaign to link the new users to                                    |
| Channels        | Text | A comma-delimited list of channel ids or names that users are invited to |

## To Do

 - [ ] Serve a customizable HTML page
 	- [ ] Serve *anything* from the plugin
 	- [ ] Store custom HTML in the configuration
 	- [ ] Allow specific channels to create guest invite links
 - [ ] Accept the user's email address
 	- [ ] Accept a POST request on a plugin page
 	- [ ] Validate the request
 	- [ ] Send a guest invitation
 	- [ ] Log the request
 - [ ] Send Lead to Marketo
 	 - https://github.com/icelander/goketo
 	 - https://developers.marketo.com/rest-api/lead-database/leads/#create_and_update

## Future Features

 - [ ] Per-channel invitation links
 	- [ ] Invitation is sent from Channel Admin
 	- [ ] Slash commands:
 		- [ ] `/invitebot setup` - Adds channel, only available to system administrators
 		- [ ] `/invitebot settings` - Opens settings dialog, only available to Channel admin or higher
 		- [ ] `/invitebot link` - Outputs the invitation link in an ephemeral message
 		- [ ] `/invitebot invite <email>` - Sends an email invitation and responds with an ephemeral message
 - [ ] Customizable fields that link to Marketo Lead Fields
 	- Use `/rest/v1/leads/describe.json` to get fields
 	- Require specific fields
 - [ ] Customizable email templates