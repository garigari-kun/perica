## perica

Github/Github Enterprise Issue to Trello Card!

## Setup

First, go get!
```
$ go get -u github.com/garigari-kun/perica
```

You need to create the config file for accessing your Trello Account and Github Account.  
Example is [this one](https://github.com/garigari-kun/perica/blob/master/example/.perica)
``` sh
$ touch $HOME/.perica
$ vi $HOME/.perica

app_key = "YOUR_TRELLO_APP_KEY"
token = "YOUR_TRELLO_APP_TOKEN"
list_id = "YOUR_TRELLO_LIST_ID"
github_token = "YOUR_GITHUB_ACCESS_TOKEN"
```

You can get ` YOUR_TRELLO_APP_KEY ` and ` YOUR_TRELLO_APP_TOKEN ` from [here](https://trello.com/app-key)

You can get ` YOUR_GITHUB_ACCESS_TOKEN ` from [here](https://github.com/settings/tokens)

## How to Use
Once your done with setup. You can use ` perica ` command like so.
```
$ perica https://github.com/garigari-kun/perica/issues/2
```
