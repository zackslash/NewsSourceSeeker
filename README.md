# News Source Seeker

News source seeker is a console application to retrieve URLS of news sources currently reporting on a given subject, it does this by querying a list of popular news search engines and scraping the URL results.

## Build


This project is written in [Google's Go programming language](https://golang.org/) (Version 1.5.1), because dependency management in Go is still in the early stages I have provided a Makefile with this project.
Building News Source Seeker should be as easy as executing the following terminal command inside the project's top directory on Linux or OSX.

`make build`

The executable 'News_Source_Seeker' and required .lst resources should now be output in a 'bin' directory inside the project dir.

### Building docker image (Optional)
If you are using **Docker** then building this project is as simple as running the following command in this project's directory

`docker build -t zackslash/news-source-seeker .`

## Run / Usage

Running this project is easy, you just need to run the executable in the bin directory followed by the subject you would like sources for

`./News_Source_Seeker <Subject name>`

## Example:

###### Command: 

`./News_Source_Seeker jack daniels`

### Run / Use through Docker (Optional)
If you are using **Docker** then after building the container, running the project is as simple as the following command.

`docker run zackslash/news-source-seeker <Subject name>`

## Example:

###### Command: 

`docker run zackslash/news-source-seeker jack daniels`

###### Example resulting JSON output:

`{"Subject":"jack daniels","Sources":["belfasttelegraph.co.uk","irishmirror.ie","breakingnews.ie"]}`

Here most 'source' results have been removed to reduce the example length. The example returns URLs of websites that have a recent article relating to 'Jack Daniels'

## Licence:
News Source Seeker is licensed under the GNU General Public License
