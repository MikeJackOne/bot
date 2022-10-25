# bot

This is a facebook messenger bot.
It has 3 functionality.
* reply to ```thank you``` message
* reply to ```generate order``` message which will generate random order information and store it into sqlite
* reply to ```get last order``` and it will send back last order information if it has.

# deploy information
Config the config.json file by specifying token(which will be used to authenticate with facebook webhook) and accessToken(your messenger portal's accessToken)
This app is deployable to heroku,just push to heroku's git and everything will take care of itself.

# demo
You can chat with this [bot](https://www.facebook.com/108071978756689/) on messenger.
