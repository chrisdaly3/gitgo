# GitGo

A minimal CLI app to query GitHub user information.
This project is intended to teach new Go users about creating simple CLI applications.

## Installation
* `go get github.com/chrisdaly3/gitgo`
* `cd $GOPATH/src/github.com/chrisdaly3/gitgo && go install`

## Usage
* Help:
```
$ gitgo
Usage: gitgo [options]
Options:
    -u, --user string
        Search Users
```


* Single User Query:
```
$ gitgo -u chrisdaly3
Searching user(s): [chrisdaly3]
Username:               chrisdaly3
Name:                   Chris Daly
Email:
Bio:            DevOps
- Backend
- GameDev
```


* Multi-User Query:
```
$ gitgo -u chrisdaly3,pmbenjamin
Searching user(s): [chrisdaly3 pmbenjamin]
Username:               chrisdaly3
Name:                   Chris Daly
Email:
Bio:            DevOps
- Backend
- GameDev

Username:        pmbenjamin
Name:            Peter Benjamin
Email:           petermbenjamin@gmail.com
Bio:             Software Engineer. Hacker. Passionate about coding, IoT, security, and teaching kids how to code.
```

### Contributing
This is mainly a teaching project for those looking to get comfortable with CLI applications in Go. Most PR's or Issues opened (if not all) will likely be declined. Feel free to fork this repo and practice with your own ideas! :) 