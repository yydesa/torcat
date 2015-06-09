A simple socks5 client program, to use as ProxyCommand in ssh, e.g.

    Host gittor
        User git
        IdentityFile ~/.ssh/hub
        ProxyCommand $HOME/bin/torcat github.com:22

to push/pull to github.

Written because there is no socks5-supporting netcat on RHEL7.

There is also tcpcat, a version that doesn't use socks5.
Noticed too late that any netcat would go for that, unlike
for 'with socks5'.
