# torcat

A simple socks5 client program, to use as ProxyCommand in ssh, e.g.

    Host gittor
        User git
        IdentityFile ~/.ssh/hub
        ProxyCommand $HOME/bin/torcat github.com:22

to push/pull to github.

Written because there is no socks5-supporting netcat on RHEL7 (unless
you install `nc`, the netcat from `nmap`).

It has one option:

`--socks host:port` changes the host and port
number of the socks port used. Default `127.0.0.1:9050`.

In putty, you need to enter the full `torcat` command
at 'Telnet command' and select 'Local', both in
the Proxy panel. (And you need to escape `\` in the
windows path as `\\` because putty.)

# tcpcat

There is also tcpcat, a version that doesn't use socks5.
Noticed too late that any netcat would go for that, unlike
for 'with socks5'.
