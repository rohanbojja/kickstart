fun main(args: Array<String>) {
    val cases = readLine()!!.toInt()
    data class Person(val times: Int, val position: Int)
    for (t in 1..cases) {
        val (n, x) = readLine()!!.split(' ').map { y -> y.toInt() }
        val arr = readLine()!!.split(' ').mapIndexed { i, y ->
            Person((y.toFloat() / x.toFloat()).toInt(), i+1)
        }
        println("${arr}")
        val arr2 = arr.sortedWith(compareBy<Person> { it.times })
        print("Case #${t}: ")
        arr.forEach { person ->
            print("${person.position} ")
        }
        println()
    }
}