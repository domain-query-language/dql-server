# DQL parser handover:

## What is done:
Fully tested tokenizer that takes in valid DQL statements and converts them into a series of tokens.
Working parser for statements and object components.
This means that we can parse the components of an object (functions, checks, handlers) and also the internal statements and expressions.
These are converted into ASTs. These ASTs do not have the ability to run, they are structs that represent the concept, with a guarantee that the data is stable and valid.
Some basic work on the top level parser was done, but it is very limited. I'd suggest rewriting this using the techniques used in the existing parsers. The tests might still be useful though,

## Where to go next:
Now that the basic compoents are working, a parser for high level DQL objects needs to be built.
Eg.
- ValueObject
- Aggregate
- CommandHandlers
- Events
- Commands
- Invariants
- Projections
- etc...

This parser would be built like the other parsers. It takes in a token stream, then parses the objects, using the object_component parser to parse the components of the objects, meaning the work is primarily focussed on parsing the unique aspects of those objects.

## How it is built:
There are three main components
Tokenizer: Turn DQL statements into a series of tokens, fail when the token doesn't make sense. Eg. A number with letters "-1sddfd".
TokenStream: Used internally in parsers, keep track of where you are in the stream, keeping the current token and the next token, makes it easy to ask questions and to navigate the stream.
Parser: Turns tokens streams into ASTs, keep going until there are no more ASTs to produce or it hits an error

The parsers rely on the token stream internally, and parsers can be created from a token stream. 
The parser are intended to be composible, a parser can be used inside another parser. Eg. The object_components parser used the statement parser internally, because functions contain statements. This was done to make it easier to build the parser from the ground up, allowing more complex parser to built from simpler components.
As each parser uses a tokenstream, it is fairly easy to do handover of token streams once a sub parser has completed it's work, you simply request the token stream from the sub parser and then set it as your own.

Each of the parser tests turn a string/strings into an AST/ASTs. The results are compared by turning the AST back into a string, and then comparing that to the string created by the hardcoded result AST. This is useful as it makes it very easy to diagnose what went wrong. This does mean that you need to write a "String" method for each AST, its a little extra work, but it really helps in diagnosing parsing issues, so it's worth the effort.

## Code structure:
All the code for this component is stored in the parser folder. Each child folder is named appropriately for what it contains.

## Extension:
Each of these components is designed to be extensible. They are based on the work on Thornsten Bell and his book on writing interpreters in GoLang. I recommend giving it a read before you start editing anything core.
If you need to parse a new kind of object, statment, component, etc.., I'd do the following
- Write the test
- Create the ASTs (so the test will actually run)
- Write the parser function in the appropriate parser