#Branchlog for Caching Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the Caching lesson](https://www.boot.dev/lessons/f2ba2b87-38fe-467c-abd7-14716f955169)
and make "moving around the map" more snappy by implementing caching.

##Commit History
###Commit C2: Fixing Debug Start
I wasn't running the debugging commands I thought I was. I've changed what I need to in order to
make this a stable commit that without syntax errors as opposed to all those caught by go vet . in
commit c1.
- Replaced types with placeholder ints but commented what things really should be.
- Added placeholder empty variables and returns to make methods syntax correct.


###Commit C1: Setup to Develop Caching
This commit just moves things around and writes down what the lesson instructions tell me to do
without making serious attempts to accomplish those instructions. I'm definitely adopting a
workflow of stablizing code, documenting it, commiting all that, and working on the next step.
Some of the moves are designed with that in mind.
- Created internal directory
- Created pokecache direcotry within internal directory
- Created pokecahce.go in pokecache directory
- Added cacheEntry struct
- Added Cache struct
- Added incomplete NewCache function
- Added empty but commented Cache.Add method
- Added empty but commented Cache.Get method
- Added empty but commented Cache.reapLoop method
- Added comments to advise direction of further development
- Renamed changelog_pokeapi.md to pokeapi.md
- Moved pokeapi.md to branchlogs directory
- Started this branchlog
