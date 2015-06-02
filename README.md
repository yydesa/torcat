A simple socks5 client program, to use as ProxyCommand in ssh, e.g.

    Host gittor
        User git
        IdentityFile ~/.ssh/hub
        ProxyCommand $HOME/bin/torcat github.com:22

to push/pull to github.
