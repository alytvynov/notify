# notify - send desktop notification when command completes

## Usage: notify [command]

Notification contains command executed, exit status and duration.

Minor imperfections due to notifications bing sent via AppleScript:
- notifications are shown from Script Editor in notification center
- clicking on a notification opens Script Editor

Known issues:
- exit code from notify doesn't match exit code from command
- common shell syntax is not supported; e.g. notify GOOS=linux ./make.bash won't work, GOOS=linux notify ./make.bash is the correct way to run it
