import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class WordsTest {

    @Test
    fun longestConsec() {
        assertEquals(
            "abigailtheta",
            longestConsec(arrayOf<String>("zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"), 2)
        )

        assertEquals(
            "oocccffuucccjjjkkkjyyyeehh",
            longestConsec(
                arrayOf<String>(
                    "ejjjjmmtthh",
                    "zxxuueeg",
                    "aanlljrrrxx",
                    "dqqqaaabbb",
                    "oocccffuucccjjjkkkjyyyeehh"
                ), 1
            )
        )
    }
}