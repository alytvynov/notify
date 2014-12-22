# notify - send desktop notification when command completes

## Usage: notify [command]

Notification contains command executed, exit status and duration.

Minor imperfections due to notifications bing sent via AppleScript:
- notifications are shown from Script Editor in notification center
- clicking on a notification opens Script Editor

Known issues:
- exit code from notify doesn't match exit code from command
