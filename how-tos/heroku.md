# Heroku

- Heroku is a container-based cloud Platform as a Service (PaaS). Developers use Heroku to deploy, manage, and scale modern apps. This platform is elegant, flexible, and easy to use, offering developers the simplest path to getting their apps to market.

- Heroku is fully managed, giving developers the freedom to focus on their core product without the distraction of maintaining servers, hardware, or infrastructure. 
- The Heroku experience provides services, tools, workflows, and polyglot support—all designed to enhance developer productivity.

# Overview

The details of Heroku's Go Support are described in the [Heroku Go Support article.](https://devcenter.heroku.com/articles/go-support "Heroku Go Support article.")
Also, we need the next features:
- Go1.7+ installed.
- $GOPATH/bin has been added to your $PATH.
- [Godep installed](https://github.com/tools/godep "Godep installed").
- [a free Heroku account](https://signup.heroku.com/ "a free Heroku account").
- [the Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli "the Heroku CLI").

####NOTICE 
Heroku Go support will be applied to applications only when the application has a file named Godeps/Godeps.json in the root directory. Even if an application has no dependencies, it must include a basic Godeps/Godeps.json file in order to be recognized as a Go application.

##Prepare your app
The next steps you can do with it:
https://devcenter.heroku.com/articles/getting-started-with-go#prepare-the-app
