# csvgen

Generate a file called _"generated.csv"_ with random values of the chosen type
> Just a tool I made to speed up testing for some stuffs I'm doing at work


## How to use it

Run the program followed by: number of rows to generate (optional, if missing then it will randomized between 1 and 100) and the types.

Example (Windows):
`$ .\csvgen.exe 1 string number pippo:pluto:paperino color`
> Possible outcome inside generated.csv:
> knKNBO7f, 32, pippo, #ff3265


For all the description about each type just run the command "help", "-help" or "--help".
> I was too ~lazy~ busy to report also here


All the inputs (help command and type identifiers) are case insensitive 
