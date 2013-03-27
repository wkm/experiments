import java.util.concurrent.{ScheduledThreadPoolExecutor, TimeUnit}


/*
 *     $ scala ScalaOutOfOrderConstructor.scala 
 *     VALUE: null
 *     done
 *     VALUE: oh hai this is cat
*/
val scheduledPool = new ScheduledThreadPoolExecutor(1)
scheduledPool.scheduleAtFixedRate (
	new Runnable {
		def run {
			println("VALUE: "+value)
		}
	},
	0, 100, TimeUnit.MILLISECONDS
)

Thread.sleep(10)
val value = "oh hai this is cat"

println("done")