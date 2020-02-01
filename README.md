# Mattermost Invitebot Plugin

## What is it?

This plugin automates the process of inviting guests to a Mattermost server. This link can be shared publicly, and users just have to enter their email addresses to get an invitation to join specific channels on a Mattermost server. This is great for onboarding new community members or engaging with conference attendees.

## Configuration Settings

| Setting         | Type | Default Value                                                            |
|:----------------|:-----|:-------------------------------------------------------------------------|
| InviteBotID     | ID   | ID of the InviteBot 														|
| UserID          | ID   | ID of the user who created the Invitebot 							    |
| LandingHTML     | Text | Stored in `assets/landing_default.html`                                  |
| ThankYouHtml    | Text | stored in `assets/thankyou_default.html`                                 |
| ErrorHTML       | Text | stored in `assets/error_default.html`                                    |
| MarketoCampaign | Text | The campaign to link the new users to                                    |
| Channels        | Text | A comma-delimited list of channel ids or names that users are invited to |

## To Do

 - [ ] Serve a customizable HTML page
 	- [x] Serve *anything* from the plugin (v0.0.1)
 	- [ ] Store custom HTML in basic configuration (v0.1.0)
 	- [ ] Allow specific channels to create guest invite links
 - [ ] Accept the user's email address (v0.1.0)
 	- [ ] Accept a POST request on a plugin page (v0.1.0)
 	- [ ] Validate the request (v0.1.0)
 		- [ ] Validate that it's an email address (v0.1.0)
 		- [ ] Check if it's an existing user (v0.1.0)
 	- [ ] Send a guest invitation (v0.1.0)
 	- [ ] Log the request (v0.1.0)
 - [ ] Send Lead to Marketo (v0.2.0)
 	 - https://github.com/icelander/goketo
 	 - https://developers.marketo.com/rest-api/lead-database/leads/#create_and_update

## Future Features

 - [ ] Multiple Invitebots
 - [ ] Bot User
 - [ ] Customizable fields that link to Marketo Lead Fields
 	- Use `/rest/v1/leads/describe.json` to get fields
 	- Require specific fields
 - [ ] Customizable email templates
 - [ ] Admin account approval