#Branchlog for Catch Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the Catch lesson](https://www.boot.dev/lessons/ed962683-cb2d-4989-99e9-5cfa144810b5)
by adding a command that allows the user to attempt to catch the pokemon they name. Personally I'd
prefer if this was locked behind finding said pokemon, but we'll save that for an after-class
feature.

##Checklists
###Checklist for Main Ready Commit
- (✓) All intended actual changes commited in current or last commit
- (✓) README.md updated to reflect change from last main branch commit
- (✓) This log updated to reflect Main Ready Commit
- (oops) Executable re-compiled
- (✓) Lesson Passed / Requests Met

###Feature Request Checklist
- (✓) Add a catch command. It takes the name of a Pokemon as an argument. Example usage:
	Pokedex > catch pikachu
	Throwing a Pokeball at pikachu...
	pikachu escaped!
	Pokedex > catch pikachu
	Throwing a Pokeball at pikachu...
	pikachu was caught!

- (✓) Be sure to print the Throwing a Pokeball at <pokemon>... message before determining if
	the Pokemon was caught or not.
- (✓) Use the Pokemon endpoint to get information about a Pokemon by name.
- (✓) Give the user a chance to catch the Pokemon using the math/rand package.
	You can use the pokemon's "base experience" to determine the chance of
	catching it. The higher the base experience, the harder it should be to catch.
- (✓) Once the Pokemon is caught, add it to the user's Pokedex. I used a map[string]Pokemon to
	keep track of caught Pokemon.
- (✓) Test the catch command manually - make sure you can actually catch a Pokemon within a
	reasonable number of tries.
- (✓) Run and submit the CLI tests.

##Commit History
###Commit H2: Main Ready (Easy Lesson)
I finished everything this needed in a couple sit downs.
- Added interfaces.go and the bodyStructure interface
- Modified callURL 's third argument to be type contrained by bodyStructure
- Added lesson-complete functionality to commandCatch
- Removed unnecessary space in argument rejection print
- Removed psuedocode and comments in commandCatch
- Modified config struct to expect a Pokedex
- Modified main function to make a Pokedex and add it to the config
- Tested Successfully
- Branchlog Updated to reflect Main Ready Commit



###Commit H1: Development Setup
I'm setting the minimum amount of code and log information to guide myself on further development.
- Added this Branchlog
- Added stable but featureless catch command
- Added comments and psuedocode hypothesizing about what needs done in commandCatch
