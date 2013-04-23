[![Build Status](https://travis-ci.org/buddhistpirate/guerilla.png?branch=master)](https://travis-ci.org/buddhistpirate/guerilla)

Guerilla is a service for generating JSON arrays of specified size. 
The content is randomly chosen from a set of source files. 
It listens on a TCP Socket for integers which are the number of lines requested.
It will pick a random file, randomly index into it and then read contigous lines until it hits the requested number of lines. If it doesn't have enough lines left in that file it will fill in the rest of the lines in the same fashion.
