### EZ install to raspberry pi

```
go get -v github.com/j4rv/cah/...
go install github.com/j4rv/cah/cah_app
```

Configure the SESSIONS_KEY env var with a random string (TODO: change this)

To execute, on a directory that has an 'expansions' folder in it:  

```
sudo -E /home/pi/go/bin/cah_app
```
