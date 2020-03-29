import java.util.Arrays;
import java.util.Comparator;

public class Words {

    public static String longestConsec(String[] words, int n) {
        if (n > words.length) {
            return "";
        }
        final String[] combine = combine(words, n);
        Arrays.sort(combine, Comparator.comparingInt(String::length).reversed());
        return combine[0];
    }

    public static String[] combine(String[] words, int n) {
        int len = words.length - n + 1;
        String[] words2 = new String[len];

        for (int i = 0; i < len; i++) {
            StringBuilder word = new StringBuilder();
            for (int j = 0; j < n; j++) {
                word.append(words[i + j]);
            }
            words2[i] = word.toString();
        }
        return words2;
    }

    public static String[][] slide(String[] list, int inGroupsOf) {
        int newSize = list.length - inGroupsOf + 1;
        String[][] res = new String[newSize][inGroupsOf];
        for (int i = 0; i < newSize; i++) {
            for (int j = 0; j < inGroupsOf; j++) {
                res[i][j] = list[i+j];
            }
        }
        return res;
    }
}
