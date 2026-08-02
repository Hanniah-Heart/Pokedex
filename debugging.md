# Debugging Log for PokeAPI Branch
This markdown file will be where I write what's going on in the commits I make while
trying to write and debug my program to add the intended feature of this branch.

I'm trying to follow the [feature requests of the PokeAPI lesson](https://www.boot.dev/lessons/813eafe1-2e1d-42a0-b358-53e0f4d4fdc8)
and add onto the initial CLI commands with commands that move back and forth between
PokeAPI location endpoint pages.

##Changelog
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
