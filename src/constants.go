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

// Prefix to remove when normalizing a URL
const webPrefix string = "www."

// This value is replaced with the actual query in seed URL strings
const queryPlaceholder string = "<QUERY>"

// Prefix that needs to be removed in each result (Google specific)
const googleURLPrefix string = "/url?q="

// Regex rule should match a valid URL beginning with http or https
const validURLRegex string = `(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`