package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

// validateDomainFormat checks if the domain matches basic domain name pattern
func validateDomainFormat(domain string) bool {
	// This pattern matches:
	// - Starts with alphanumeric or hyphen
	// - Contains alphanumeric, hyphens (not consecutive)
	// - Contains at least one dot
	// - Ends with a valid TLD (2 or more chars)
	// - No consecutive dots
	pattern := `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](\.[a-zA-Z]{2,})+$`

	matched, err := regexp.MatchString(pattern, domain)
	if err != nil {
		log.Printf("ERROR: Domain validation regex error: %v", err)
		return false
	}
	return matched
}

func CheckDomain(domain string) {
	log.Printf("INFO: Starting validation for domain: %s", domain)

	// Validate domain format first
	if !validateDomainFormat(domain) {
		log.Printf("ERROR: Invalid domain format: %s", domain)
		fmt.Printf("Invalid domain format. Please enter a valid domain (e.g., example.com)\n")
		return
	}

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	fmt.Println("-----------------------DOMAIN RELATED INFORMATION---------------------------")
	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	} else if err != nil {
		log.Printf("Error: %v\n", err)
	}

	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, txt := range txtRecords {
			if strings.HasPrefix(txt, "v=spf1") {
				hasSPF = true
				spfRecord = txt
			}
		}
	} else {
		log.Printf("Error: %v\n", err)
	}

	dmarcDomain := "_dmarc." + domain
	dmarcTxtRecords, err := net.LookupTXT(dmarcDomain)
	if err == nil {
		for _, txt := range dmarcTxtRecords {
			if strings.HasPrefix(txt, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = txt
			}
		}
	} else {
		log.Printf("Error: %v\n", err)
	}

	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX Record: %t\n", hasMX)
	fmt.Printf("Has SPF Record: %t\n", hasSPF)
	if hasSPF {
		fmt.Printf("SPF Record: %s\n", spfRecord)
	}
	fmt.Printf("Has DMARC Record: %t\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("DMARC Record: %s\n", dmarcRecord)
	}
	fmt.Println("---------TYPE NEXT DOMAIN ADDRESS TO CONTINUE OR QUIT TO EXIT----------")
}
