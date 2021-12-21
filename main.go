package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/cloudflare"
	"github.com/libdns/dnspod"
	"github.com/libdns/libdns"
)

type Provider interface {
	libdns.RecordGetter
	libdns.RecordAppender
	libdns.RecordSetter
	libdns.RecordDeleter
}

func main() {
	domain := flag.String("domain", "", "Domain")
	dnsProvider := flag.String("provider", "", "DNS Provider")
	dnsApiToken := flag.String("token", "", "DNS Provider API Token")
	acmeChallengeType := flag.String("challenge-type", "", "ACME challenge type to use")

	flag.Parse()

	provider := getProvider(*dnsProvider, *dnsApiToken)

	if *acmeChallengeType == "dns-01" {
		certmagic.DefaultACME.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: provider,
		}
	}

	ctx := context.Background()

	fmt.Println("Getting cert for domain: " + *domain)

	err := certmagic.ManageSync(ctx, []string{*domain})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getProvider(providerId, dnsApiToken string) Provider {
	switch providerId {
	case "cloudflare":
		return &cloudflare.Provider{
			APIToken: dnsApiToken,
		}
	case "dnspod":
		return &dnspod.Provider{
			APIToken: dnsApiToken,
		}
	default:
		return nil
	}
}
