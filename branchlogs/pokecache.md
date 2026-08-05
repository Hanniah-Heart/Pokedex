#Branchlog for Caching Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the Caching lesson](https://www.boot.dev/lessons/f2ba2b87-38fe-467c-abd7-14716f955169)
and make "moving around the map" more snappy by implementing caching.

##Commit History
###Commit C7: Encapsulated mapOut Function
This commit we made sure that mapOut actually implements the Cache.Add and Cache.Get methods. We
also fixed some mistakes in new caches and made sure all methods refer to the Cache address not a
copy of the cache.
- Updated this log
- Modified mapOut function use encapsulatin methods
- Modified NewCache to initialize properly
- Modified Cache struct to use a mutex rather than nonsensically using a pointer to a mutex
- Modified all Cache methods to refer to the address of the cache, not a copy


###Commit C6: Stable but Unencapsulated mapOut Update
This commit I was trying to update the mapOut function check if there was already a cached entry 
for the information we're collecting from the map and return the cache instead. I got that working
and stable, but I realized when I was finally done that this was obviously supposed to implement
the Cache.Get method. I did not do that. Next Commit I will probably switch to doing that.
####Current Goal:
Update your code that makes requests to the PokeAPI to use the cache. Create the cache once and 
reuse it in your PokeAPI request layer. If you already have the data for a given URL (which is our 
cache key) in the cache, you should use that instead of making a new request. Whenever you do make 
a request, you should add the response to the cache.
####Changes
- Updated this log.
- Modified the mapOut funciton to check for cached information and call it instead.
- Modified go.mod to capitalize the url as appropriate to my github repository url.
- Modified pokecache.go to export variables as needed.
- Modified main.go to accept exported variables as needed.
- Modified structs.go to export fields as needed.



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
