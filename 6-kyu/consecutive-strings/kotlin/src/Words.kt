import java.lang.StringBuilder

fun message(): String {
    return "Have a fine day!"
}

fun longestConsec(words: Array<String>, n: Int): String {
    var largest = ""

    for (i in 0..words.size - n step 1) {
        var builder = StringBuilder()
        for (j in i until (i + n)) {
            builder.append(words[j])
        }

        if (builder.length > largest.length) {
            largest = builder.toString()
        }
    }
    return largest
}


fun main() {

//    val names = Array<String>(3) { "Benjamin"; "Pénélope", "Ewan", "Dorian"}
    val names = arrayOf("Benjamin", "Pénélope", "Ewan", "Dorian")
    for (name in names) {
        println("Hello $name")
    }

    names.iterator().forEach {
        println("Hello, $it (message is: '${message()}')")
    }

    val sequence = generateSequence(1) { it + 1 }

    val windows = sequence.windowed(size = 5, step = 1)
    println(windows.take(4).toList())

    val combined = listOf("Benjamin", "Pénélope", "Ewan", "Dorian")
        .windowed(size = 2)
        .toList()

    println(
        combined
    )

    println(
        listOf("Benjamin", "Pénélope", "Ewan", "Dorian")
            .windowed(step = 2, size = 1)
            .toList()
    )

}