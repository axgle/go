Based on https://github.com/mattn/go-mruby

	package main

	import "github.com/mattn/go-mruby"

	func main() {
		mrb := mruby.New()
		defer mrb.Close()	
		mrb.Run("puts 'hello world' ")
	 
	}

