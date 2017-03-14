#MinimalCI

CI/CD need not be complicated, and minimalCI aims to create no-nonsense CI for any github project.  Creating a CI process is as easy as a single command.

##About MinimalCI

Distilled to its essence a CI/CD process involves doing something when a change is introduced into a code repository.  The main requirement is a component to watch a repository and execute something when code changes.  Working with other frameworks like Jenkins that contain a massive amount of technical debt, a large learning curve, unnecessary security, and other instabiities can be counter productive.  MinimalCI watches your repository or listens for events and then executes the command specified.  How and what gets run is up to you.

As with the more modern pipeline CI frameworks your build, test, deployment lives in the git repo itself.  MinimalCI uses a file called minimalCI.yaml that contains the starting point for the 

MinimalCI is written and go and is the beneficiary of many contributions to the open source community

##Getting Started:

##To install minimalCI:
		sudo bash https//minimalCI.jfelten.ghpages.org/installMinimalCI.sh

####OR
		go get minimalCI

###To poll a public github repo for PRs

		minimalCI poll_github <GITHUB_URL> --cron <CRON_EXP> --githubPR <

###To create a github webhook listener for PRs:

		minimalCI webhook_github <GITHUB_URL> --githubPR --key <KEY>

###To poll for checkins to a vanilla git repo

		minimalCI poll_git <CRON_EXP> --git <GIT_URL>
That's it.

###Other Examples

##The minimalCI configuration file

MinimalCI will look for the optional minimalCI.yaml file in the root of a project.  If found it will execute the command specified.  If no command is specified or no file is found minimalCI will 
