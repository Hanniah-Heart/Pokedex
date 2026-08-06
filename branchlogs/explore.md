#Branchlog for Explore Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the Explore lesson](https://www.boot.dev/lessons/e53abbb4-5d8a-4feb-ba08-828f03311e51)
by adding a command that allows the user to collect information about specific locations from
the map. This will implement positional arguments.

##Commit History
###Commit E4: Command Refactor and Explore Caching
There was previously a block of code I had basically repeated from mapOut to commandExplore. I've
refactored that block into its own function to dry out the code. I've also condensed several error
checks into one line with the call that risked their error, thereby psuedo-drying 3 lines per
several calls. The code runs beautifully and is much more concise. I'm so happy.
- Updated this branchlog
- Modified commands.go functions to be line-concise and DRY.
- Added functionality to commands.go which ensures DRY implementation of caching and http calls.


###Commit E3: Working Explore Command
I already have the explore command working correctly. Maybe there's somethings to clean up before
main ready, but this looks nice. I also corrected the handling of arguments. Instead of expecting
nil as the zero value of an empty slice like I was (which make sense but I still should have
tested), I'm now checking the length of the arguments slice.
- Updated this Branchlog
- Added locationAreaStructs.go to properly unmarshall the data commandExplore collects.
- Modified commandExplore to expected functionality
- Modified all commands to handle argument count more appropriately


###Commit E2: Stable Argument Rejection
This commit I just made sure that all commands can handle arguments passed to them even though the
default behavior I've given them all is to say that arguments aren't allowed. This will change but
is syntax-error free progress versus said functions not even listening for any user arguments.
- Updated this Branchlog
- Modified main.go to pass arguments through the cmd.callback
- Modified commands.go to include stable code handling arguments from the cmd.callback
- Modified cliCommand struct to expect a list of strings as well


###Commit E1: Development Setup
I'm setting the minimum amount of code and log information to guide myself on further development.
- Added this Branchlog
- Modified commandList to include explore
- Added stable but otherwise empty commandExplore function
- Added code and comment to main.go to prepare to implement arguments into cmd.callback

##Checklist for Main Ready Commit
- ( ) All intended actual changes commited in last commit
- ( ) README.md updated to reflect change from last main branch commit
- ( ) This log updated to reflect Main Ready Commit
- ( ) Executable re-compiled
