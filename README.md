<p align="center"> <b>Reverse Proxy</p>
<p align="center">
  <img src="https://github.com/gh-ninja/reverse-proxy/blob/main/s.png" />
</p>

# Dependencies
- [Go](https://golang.org/)
- [Make](https://www.google.com/search?q=how+to+install+and+use+make&sxsrf=AOaemvIDurREKHM1mD434VRIASsFjPm0lQ%3A1630461057089&source=hp&ei=gdwuYZKFA4Ou5OUPsp-akAo&iflsig=ALs-wAMAAAAAYS7qkcG54HJBz1ihgC8eVKHYVKqP4gKO&oq=how+to+install+and+use+make&gs_lcp=Cgdnd3Mtd2l6EAMyBQgAEMsBOgcIIxDqAhAnOgQIIxAnOgsIABCABBCxAxCDAToOCC4QgAQQsQMQxwEQowI6EQguEIAEELEDEIMBEMcBEKMCOgsILhCABBCxAxCDAToRCC4QgAQQsQMQgwEQxwEQ0QM6BQguEIAEOgoIABCxAxCDARBDOgQIABBDOgYIIxAnEBM6CAgAEIAEELEDOggILhCxAxCDAToOCC4QsQMQgwEQxwEQowI6CAguEIAEELEDOgUIABCABDoGCAAQFhAeOgcIIxCwAhAnOgQIABANOgYIABANEB5QxgxY7ntgi5ABaBBwAHgCgAG9AogBvzWSAQgwLjM0LjcuMZgBAKABAbABCg&sclient=gws-wiz&ved=0ahUKEwjSyYDh1NzyAhUDF7kGHbKPBqIQ4dUDCAg&uact=5)
# Suport
- [x] Termux (android/afsd kernel)
- [x] linux (kernel)

# Install:
###### Termux:
1 step: Install Go-lang, Git and Make `pkg install make -y && pkg install golang -y && pkg install git -y`<br>
2 step: Clone Rp: `git clone https://github.com/gh-ninja/reverse-proxy.git`<br>
3 step: Build Tool: `make build` <br>

###### Linux (deb):
1 step: Install Go-lang, Git and Make `apt install make -y && apt install golang -y && apt install git -y`<br>
2 step: Clone Rp: `git clone https://github.com/gh-ninja/reverse-proxy.git`<br>
3 step: Build Tool: `make build` <br>

# Args:
- listen: `--listen, -l` <br>
- set: `--set , -s` <br>
- localconfig: `--localconfig, -lc` <br>


###### example
- 1 listen using local config: `./reverse-proxy --listen --localconfig` or `./reverse-proxy -l -lc`<br>
- 2 listen using set config: `./reverse-proxy --listen --set` or `./reverse-proxy -l -s`<br>

##### Edit Local Config
- Tip: _To change the default settings, edit the file: [config.json](https://github.com/gh-ninja/reverse-proxy/blob/main/config.json)_


##### Config
- Target: target url where reverse proxy will redirect requests<br>
- Kepalive: enable keepalive connection, set `true` for enabled and `false` for disabled<br>
- web-interface-port: port where the web interface is running<br>
- proxy-port: port where the reverse proxy will run <br>
----
<p align="center">🚀 by Ninja</p>
