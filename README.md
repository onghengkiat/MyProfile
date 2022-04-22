# My Profile
This is a mini project created by me to practice on the javascript, html and css.
Also it shows up about myself.

## Technologies
---
1) Golang
2) Javascript
3) Bootstrap
4) HTML and CSS
5) Docker

## Scope of Functionalities
---
- View certain information about myself

## User Guide
---
The website is hosted on <a class="font-weight-bolder" href="https://onghengkiatv1.herokuapp.com/">My Profile</a>.

## Developer Guide

### Testing on localhost without using docker: 

1) go run main.go
</br>

The website is hosted on heroku using docker container following this 
<a href="https://dev.to/erenaspire7/deploying-a-dockerized-flask-app-to-heroku-5h7j">guide</a>.

### Testing on localhost using docker: 

1) Run "docker build . -t (image_name)"
2) Run "docker run -it --name (container_name) -p 8080:8080 (image_name)"

**Notes: Make sure your working directory is at MyProfile**

### Deploy

1) Run "heroku login". If heroku login shows browser cant be opened, then try "heroku login -i"
2) Run "heroku create (app-name)". Remember the Heroku git URL.
3) Run "git remote add (branch-name) (git-URL)"
4) Run "git add ."
5) Run "git commit -m "First Commit"
6) Run "heroku stack:set container"
7) Run "git push (branch-name)"

**Notes: Steps 2, 3 and 6 are for first deployment only**

