# Gonginx
Gonginx is an Nginx configuration parser helps you to parse, edit, regenerate your nginx config files in your go applications. It makes managing your balancer configurations easier. 



Note: This project fork from tufanbarisyildirim/gonginx and add some awesome features/bugfixes 🎉!


# Feature  


- Parse \*\_by\_lua\_block code block simply  
- More rubust code,will not panic

# Architecture  
## Core Components  


- ### [Parser](/parser) 
  Parser is the main package that analyzes and turns nginx structred files into objects. It basically has 3 libraries, `lexer` explodes it into `token`s and `parser` turns tokens into config objects which are in their own package, 
- ### [Config](/config.go)
  Config package is representation of any context, directive or their parameters in golang. So basically they are models and also AST
- ### [Dumper](/dumper.go)
  Dumper is the package that holds styling configuration only. 

# Examples
- [Formatting](/examples/formatting/main.go)
- [Adding a Server to upstream block](/examples/adding-server/main.go)

# TODO



- [ ]  associate comments with config objects to print them on config generation and make it configurable with `dumper.Style`
- [ ]  move any context wrapper into their own file (remove from parser)
- [ ]  Parse included files recusively, keep relative path on load, save all in a related structure and make that optional in dumper.Style
- [ ]  Implement specific searches, like finding servers by server_name (domain) or any upstream by target etc.
- [ ]  more exampels
- [ ]  link the parent directive to any directive for easier manipulation



# Thanks

Fork from [tufanbarisyildirim/gonginx](https://github.com/tufanbarisyildirim/gonginx) , thanks these guys who made that awesome project!



# License

[MIT License](LICENSE)