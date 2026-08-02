# Debugging Log for PokeAPI Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the PokeAPI lesson](https://www.boot.dev/lessons/813eafe1-2e1d-42a0-b358-53e0f4d4fdc8)
and add onto the initial CLI commands with commands that move back and forth between
PokeAPI location endpoint pages.

##Changelog
###Commit P5: Refactor 1
Here we clean up what we're doing a little bit. I don't want to scroll endlessly on main.go this
time around. I've moved the structs and commands into their own files. I still think repl.go is
poorly named for what little it has, but we'll worry about that later. 
- Modified main.go to remove the debug string.
- Refactored main.go into commands.go and structs.go to limit file length while my IDE 
	is nano and doesn't have code folding.


###Commit P4: Pointer to Config
We're successfully passing a config into the commands now.
- Modified main.go to pass map_conf into the commands rather than the debug string.

###Commit P3: Initialize a Config
The goal we're working towards is passing a pointer to a config into a command and changing
the original config repeatedly with commands. We're not there yet. Here we're just initializing
the config that we'll be changing. In this case that's map_conf. Since we want to see that it
worked and go will error if we don't do anything with it, we simply print it immediately.
- Modified main.go to initialize a map_conf config and then print it immediately.

###Commit P2: Pointer to String
We're now handling the thing passed into the commands as a reference to a string rather than a copy
of the string. We also demonstrate that this is consistently changing the string throughout the
entire use of the program by appending the string each time help is used. We don't do this with
exit since exit will only run once and the contrast is helpful for debugging.
- Modified main.go command functions to handle a pointer to a string rather than a copy.


###Commit P1: Debug String
My goal here was simply to successfully pass SOMETHING into one of these commands.
In go you have to do something with any variable you call and I printing strings helps,
so I just printed a string that helps me know the code is reading that string value.
- Modified main.go to pass a string into each command and call that string.
