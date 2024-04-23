package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("domain,hasMX,hasSPF,sprRecords,hasDMARC,dmarcRecord\n\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			hasSPF = true
			spfRecords = txtRecord
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	}

	for _, dmarcRecord = range dmarcRecords {
		if strings.HasPrefix(dmarcRecord, "v=DMARC1") {
			hasDMARC = true
			break
		}
	}

	fmt.Printf("%s,%t,%t,%s,%t,%s\n", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord)
}
