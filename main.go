package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gocolly/colly"
)

/*
TODO: define a struct to contain the product data and be sent to the MongoDB storage
~ 7/24
*/

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("/c/Users/안상태/go/src/github.com/ahnsv/colly/.aws/credentials.csv", "ahnsv"),
	})
	svc := dynamodb.New(sess)
	// Test connection
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("url_name"),
	}
	result, err := svc.DescribeTable(req)
	if err != nil {
		fmt.Printf("%s", err)
	}
	table := result.Table
	fmt.Printf("done", table)
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: birchbox.com
		colly.AllowedDomains("www.birchbox.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start scraping on birchbox's best sellers
	c.Visit("https://www.birchbox.com/category/149")
}
