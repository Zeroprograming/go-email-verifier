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
	fmt.Print("domain,hasMX,hasSPF,sprRecord,hasCNAME,hasTXT\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n ", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasCNAME, hasDMARC, hasTXT bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	for _, mxRecord := range mxRecords {
		if strings.HasSuffix(mxRecord.Host, domain) {
			hasMX = true
			break
		}
	}

	if err == nil {
		hasMX = true
		log.Printf("Error hasMX: %v\n", err)
	}

	txtRecords, err := net.LookupTXT(domain)

	if err == nil {
		hasTXT = true
		log.Printf("Error hasTXT: %v\n", err)
	}

	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			hasSPF = true
			spfRecord = txtRecord
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err == nil {
		hasDMARC = true
		log.Printf("Error hasDMARC: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, spfRecord, hasCNAME, hasTXT, hasDMARC, hasDMARC, dmarcRecord) // Corregir esta línea
	fmt.Println("Domain: ", domain)

}
