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

import "bufio"
import "os"
import "strings"

// reads a whole file into memory and returns a slice of its lines.
// Ignores comments and blank lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	
	if err != nil {
		return nil, err
	}
	
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		newLine := scanner.Text();
		
		if(!strings.HasPrefix(newLine, "#") && newLine != "") {
			lines = append(lines, newLine)
		}
	}
	return lines, scanner.Err()
}