
trait Configuration {
}

trait WithConfig {
	def config: Configuration
	def execute()
}

class Foo extends WithConfig {
	object config extends Configuration {
		var int = 12
		var str = "oh hai"
	}

	def execute() {
		println("oh hai")
		println("int: "+(config.int))
		println("str: "+(config.str))
	}		
}

val instance = new Foo()
instance.config.int = 143
instance.config.str = "kitty kat"

val newInstance = new Foo()
newInstance.config.int = 314159
println("c: "+instance.config.int)
println("n: "+newInstance.config.int)