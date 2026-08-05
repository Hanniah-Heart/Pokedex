#Branchlog for Caching Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the Caching lesson](https://www.boot.dev/lessons/f2ba2b87-38fe-467c-abd7-14716f955169)
and make "moving around the map" more snappy by implementing caching.

##Commit History
###Commit C5: Stable reapLoop Method
We now have a stable Cache.reapLoop method that's called by Cache.NewCache on a separate goroutine
to delete cache entries that are older than the interval. The interval is also now a setting in the
main config.
- Updated this log
- Modified reapLoop to delete old entries on a ticker
- Modified NewCache to call reapLoop on a new goroutine
- Modified the config struct to include the reaping interval
- Modified the map_conf declaration to include the reaping interval
- Added comments to reflect possible future features
	- Runtime configurable reaping interval
	- Cache entry refresh upon get call
- Built and tested the program to make sure it's still working
	- This likely is irrelevant since the main program doesn't call the pokecache package but


###Commit C4: Stable Get Method
Like the last method, I don't have confirmation this is functional but I have confirmation it
is error free. All I did is add code that seems like it should handle Cache.Get 's purpose
correctly.
- Updated this log
- Added parser-error free code to the Cache.Get method
- Corrected this log's Commit C3 section to say parser-error free rather than error free


###Commit C3: Stable Add Method
I have no errors here and my types are all correct again and we've added what might be the correct
way of handling the Cache.Add method. I don't have tests for it, but it's stable and progress.
- Updated this log
- Added parser-error free code to the Cache.Add method
- Included appropriate mutual exclusions inside Cache.Add and Cache.Get
- Fixed Types to how they should be
- Modified NewCache to declare a Cache without any value assignment
- Modified NewCache to return the address of the new cache


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
