certgrabber is a thin wrapper around [certmagic][0]. It can be used to get
TLS certificates using the ACME protocol (ie LetsEncrypt, etc), similar to
certbot.

The main advantage of certgrabber is that it stores certs in the same location
and format as Caddy, [boringproxy][1], and any other tools based on certmagic.

This can be useful if you have a tool based on certmagic that doesn't support
the DNS challenge, which is currently the case for boringproxy. It lets you
get certs using the DNS challenge with certgrabber, and the certs can then
be used by boringproxy.

[0]: https://github.com/caddyserver/certmagic

[1]: https://github.com/boringproxy/boringproxy
