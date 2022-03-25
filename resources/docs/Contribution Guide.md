{.text-5xl}
# Contribution Guide

{.mt-5}
I should assume you've successfully installed your go in your machine, to start working with this, you should fork a copy of `master` branch to your github, checkout under your `$GOPATH/src/` folder.

{.my-5}
If you want to quickly try Gorvel, please follow bellow source, make sure your port `8080` is open to serve your local http, or modify your gorvel `.env` file.

```bash
$> echo $GOPATH
/Users/johndoe/go
$> cd /Users/johndoe/go
$> mkdir src/
$> wget -c https://github.com/daison12006013/gorvel/archive/refs/heads/master.tar.gz -O - | tar -xz
$> cd gorvel-master
$> bash ./serve
```

{.text-3xl .mt-5}
## For Security Issues

{.mt-5}
If you found any security issues, things you may find that breaks could someone bypass Gorvel nor Go's security, please send an email to `daison12006013@gmail.com`
