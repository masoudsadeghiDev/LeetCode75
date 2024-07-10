class Solution {
    public String mergeAlternately(String word1, String word2) {
        StringBuilder result = new StringBuilder();
        boolean toggle = true;
        int counter = 0;
        int resultL = word1.length() + word2.length();
        int pt1 = 0;
        int pt2 = 0;

        while (counter < resultL) {
            if (toggle) {
                if (pt1 < word1.length()) {
                    result.append(word1.charAt(pt1));
                    pt1++;
                    counter++;
                }
            } else {
                if (pt2 < word2.length()) {
                    result.append(word2.charAt(pt2));
                    pt2++;
                    counter++;
                }
            }
            toggle = !toggle;
        }
        return result.toString();
    }
}