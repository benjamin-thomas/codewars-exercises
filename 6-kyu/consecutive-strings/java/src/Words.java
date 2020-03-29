public class Words {

    public static String longestConsec(String[] words, int n) {
        System.out.println(n);
        return words[0];
    }

    public static String[] combine(String[] words, int n) {
        int len = words.length - n + 1;
        String[] words2 = new String[len];

        for (int i = 0; i < len; i++) {
            String word = "";
            for (int j = 0; j < n; j++) {
                word += words[i + j];
            }
            System.out.println("word =" + word);
            words2[i] = word;
        }
        return words2;
    }
}
