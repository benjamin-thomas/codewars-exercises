import org.junit.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class WordsTest {

    @Test
    public void slidingWindow() {

        assertArrayEquals(new String[][]{
                {"a", "b", "c"},
                {"b", "c", "d"},
                {"c", "d", "e"},
                {"d", "e", "f"}
        }, Words.slide(new String[]{"a", "b", "c", "d", "e", "f"}, 3));

        assertArrayEquals(new String[][]{
                {"a", "b"},
                {"b", "c"},
                {"c", "d"},
                {"d", "e"},
                {"e", "f"},
        }, Words.slide(new String[]{"a", "b", "c", "d", "e", "f"}, 2));

        assertArrayEquals(new String[][]{
                {"a"},
                {"b"},
                {"c"},
                {"d"},
                {"e"},
                {"f"},
        }, Words.slide(new String[]{"a", "b", "c", "d", "e", "f"}, 1));

        assertArrayEquals(new String[][]{
                {"a", "b", "c"},
                {"b", "c", "d"},
                {"c", "d", "e"},
        }, Words.slide(new String[]{"a", "b", "c", "d", "e"}, 3));
    }

    @Test
    public void combine() {
        // noinspection SpellCheckingInspection
        assertArrayEquals(
                List.of("aabb", "bbcc", "ccdd").toArray(String[]::new),
                Words.combine(List.of("aa", "bb", "cc", "dd").toArray(String[]::new), 2)
        );

        // noinspection SpellCheckingInspection
        assertArrayEquals(
                List.of("aabb", "bbcc").toArray(String[]::new),
                Words.combine(List.of("aa", "bb", "cc").toArray(String[]::new), 2)
        );

        // noinspection SpellCheckingInspection
        assertArrayEquals(
                List.of("aabbcc", "bbccdd").toArray(String[]::new),
                Words.combine(List.of("aa", "bb", "cc", "dd").toArray(String[]::new), 3)
        );

        // noinspection SpellCheckingInspection
        assertArrayEquals(
                List.of("aabbcc").toArray(String[]::new),
                Words.combine(List.of("aa", "bb", "cc").toArray(String[]::new), 3)
        );


    }

    @Test
    public void longestConsec() {
        assertEquals("abigailtheta",
                Words.longestConsec(new String[]{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"}, 2)
        );
        assertEquals("oocccffuucccjjjkkkjyyyeehh",
                Words.longestConsec(new String[] {"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1)
        );
        assertEquals("",
                Words.longestConsec(new String[] {"zone", "abigail", "theta", "form", "libe", "zas"}, -2)
        );
        assertEquals("",
                Words.longestConsec(new String[] {"it","wkppv","ixoyx", "3452", "zzzzzzzzzzzz"}, 15)
        );
    }
}