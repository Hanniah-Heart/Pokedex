# Pokedex
Note: this README will only be updated for main branch commits or other major commits.
See the branch log by the associated branch name for more specific documentation.
## Introduction
This is a guided project instructed by [Boot.dev](https://www.boot.dev).

The opening line of the guided project is:
We're going to build a Pokedex in a command-line REPL. We'll use the incredible PokéAPI to fetch all of the data we'll need using GET requests. If you're not familiar with Pokemon, or a Pokedex, that's okay! A Pokedex is just a make-believe device that lets us look up information about Pokemon - things like their name, type, and stats.

In the second lesson we're told to keep a git location for this project without in-course instructions since we should already have compeleted the Git course earlier in our learning path. Said git course is actually the reason behind ["my webflyx project"](https://github.com/Hanniah-Heart/webflyx). I have jumped ahead to this particular project in excitement to make my first program that I can properly share with non-programmers. Before this I have officially completed 258 lessons on boot.dev (more if you count the ones I did on my ubuntu shell without logging them on the site while I was debating paying for boot.dev membership), and officially completed 2 boot.dev courses ("Learn Object Oriented Programming in Python" and "Learn Git" --the latter resulting in the webflyx repository), and 1 other project ("Build a BookBot in Python") specifically considered a boot.dev guided project. This does not count me unofficially completing the "Learn Linux" course nor the amount of the "Learn Python for Beginners" or "Learn Go" courses that I have done without finishing said courses. I'm working on this in Go as a matter of personal ambition and demonstration to myself of what my continual learning will give me with the right tools. I'd have preferred to use the Python language that most of this path has focused on so far, but, sadly, Python is an interpretted language and the methods to try to turn its programs into executables anyway... require more care and discernment before deciding to use them. I will also admit that I in fact do not know much about Pokemon and don't play it much. If you find this Pokedex good anyway, that's awesome and I hope it does good things for you. Either way thank you for trying my program and I hope you have wonderful day!

## Functionality at present
This REPL program now supports 4 recognized commands (seen under help command), and handles
errors, piped inputs with End-of-File's, and empty user input. The most relevant commands are
map and mapb which paginate through lists of 20 location names. Anything collected from the
internet is now cached temporarily and reused if possible to avoid unnecessary internet calls.


## Changelog
### Commit 6: Caching
Implements caching so that anything collected from the internet is now cached temporarily and
reused if possible to avoid unncessary internet calls.
- Updated README.md 's branchlog note to not rely on markdown that github doesn't recognize.
- Added a branchlog for the Caching feature branch
- Modified the mapOut function to implement caching
- Corrected the go.mod file 's capitalization
- Added internal/pokecache/pokecache.go which focuses on implementing caching
- Modified main.go properly intialized the modified code, data, and variables
- Modified structs.go so the config struct supports caching
- Added requested pokecache_test.go file


### Commit 5: PokeAPI
Adds the map and mapb commands for the REPL which pull information from the PokeAPI and then lists
a page of 20 locations which can be cycled through with map (forward) and mapb (back).
- Refactored main.go commands and structs into relative .go files to minimize main.go length.
- Added map command which shows the next page of locations if not saying you're on the last page.
- Added mapb command which shows the previous page of locations if not saying you're on the 1st page.
- Added a branch specific changelog describing how this commit was created from Commit 4.
#### Functionality
This REPL program now supports 4 recognized commands (seen under help command), and handles
errors, piped inputs with End-of-File's, and empty user input. The most relevant commands are
map and mapb which paginate through lists of 20 location names.


### Commit 4: Initial Commands for REPL
Replaces the generic printback with a structure that recognizes specific commands and acts on them and can handle errors.
- Modified main.go to stop printing back the first word of the input.
- Modified main.go to establish the input scanner before entering the loop.
- Modified main.go to a callable map[string]struct which defines the recognized commands (help, exit).
- Modified main.go to handle Ends of File incase of limited piped input as opposed to user input.
- Modified main.go to handle empty user input. 
- Modified main.go to handle errors appropriately even though none are expected.
#### Comment
I really want to refactor this so that there's not such a mess directly in main.go,
but I'm going to try and complete the guided instructions of this guided project before doing that.
#### Functionality
This REPL program now supports 2 recognized commands (seen under help command), and handles
errors, piped inputs with End-of-File's, and empty user input.



### Commit 3: Minimal REPL
The program is now technically a REPL. It doesn't do anything but a silly print back, but it works.
- Modified main.go to Read Evaluate Print and Loop properly instead of simply testing cleanInput.
- Added repl.log to .gitignore. That's a testing output log I'm generating manually, not actual program parts.
- Cleaned repl.go of debugging prints.
#### Functionality
The program now prompts "Pokedex > " and then takes input and spits back out "Your command was: first word " and that's it.



### Commit 2: Clean Input and Unit Testing
Completing a setup that allows the command go test ./... to confirm whether or not I'm creating an input cleaning function correctly, and making said function strip any text down to only words fielded by whitespace.
- Created the repl.go function "Clean Input" for the purpose described above.
- Created the repl_test.go file to check if Clean Input works correctly.
- Modified main.go to run repl.go 's cleanInput command on a "Hello, World!" string
- Added .gitignore with a line for *.swp to avoid committing said generated files that seem to be garbage anyway.
- Updated this changelog including correcting Commit 1, Itemized list item 2.
- Added the below functionality section.
#### Functionality
At the moment, the program still fully relies on interaction with a CLI running go commands.
It doesn't even take new strings for the cleanInput function.
You can run "go test ./..." in CLI to ask the testing import about about errors in the cleanInput function.
You can run "go run ." in CLI to which just prints "Hello, World!" and then runs cleanInput specifically on the line "Hello, World!", not a custom string.
Debugging prints are still here as this is not intended as true release version.



### Commit 1: Hello World
Creating initial files just to have any valid "Hello World" program and related documentation started.
- Added file: "go.mod" as part of infrastructure to run a Go program.
- Added file: "main.go" solely with instructions to print "Hello, World!" to demonstrate that I have a working program.
- Compiled: "pokedexcli" based on the above.
- Added this README.md to explicitly document my work.
