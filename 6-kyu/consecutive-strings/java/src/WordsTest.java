import org.junit.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

public class WordsTest {

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
}