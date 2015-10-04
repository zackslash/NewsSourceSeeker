/*
	This file is part of News Source Seeker.
	
	Copyright (C) 2015  Luke Hines

    News Source Seeker is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    News Source Seeker is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with News Source Seeker.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	
	"github.com/calbucci/go-htmlparser"
)

func main() {
	argCount := len(os.Args)

	if(argCount < 2) {
		fmt.Printf("\nPlease specify subject as parameter 1 Example: ./News_Source_Seeker Apple\n\n")
		os.Exit(0)
	}
	
	searchURLs,err := readLines("seeds.lst")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	
	blackList,err := readLines("blacklist.lst")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	
	urlRegex := regexp.MustCompile(validURLRegex)
	
	// Format query
	query := parseArgQuery();

	links := []string{}
	for _,searchURL := range searchURLs {
		//Get domain level
		seedDomain, _ := url.Parse(searchURL)
		seedRootHost := strings.Replace(seedDomain.Host, webPrefix, "", 1)

		searchURL = strings.Replace(searchURL, queryPlaceholder, query, 1)
		resp, err := http.Get(searchURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		parser := htmlparser.NewParser(string(body))
		parser.Parse(nil, func(e *htmlparser.HtmlElement, isEmpty bool) {
			if e.TagName == "a" {
				link,_ := e.GetAttributeValue("href")
	
				// Remove google link prefix
				link = strings.Replace(link, googleURLPrefix, "", 1)
		
				// Validate URL format
				validURL := urlRegex.MatchString(link)
				resultDomain, _ := url.Parse(link)
		
				// Validate URL is not a subdomain of the seed
				subdomainOfSeed := strings.Contains(resultDomain.Host,seedRootHost)
		
				// Normalise URL
				resultNewsSource := strings.Replace(resultDomain.Host, webPrefix, "", 1) 
		
				// Validate URL is not blacklisted
				blacklisted := sliceContains(blackList, resultNewsSource)
		
				if(resultNewsSource != "" && !sliceContains(links, resultNewsSource) && validURL && !subdomainOfSeed && !blacklisted) {
					links = append(links, resultNewsSource)
				}
			}
		}, nil)
	}
	
	fmt.Println(generateJSONResponse(query, links))
}
