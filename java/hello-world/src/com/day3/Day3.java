import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day3 {
    public static void main(String[] args) {
        try {
            String content = Files.readString(Paths.get("../../", "data/day3/data.txt"),
                    StandardCharsets.UTF_8);
            Day3 day3 = new Day3();
            int result = day3.getMulsResult(content);
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        }

    }

    public ArrayList<String> getMulsFromText(String text) {
        ArrayList<String> muls = new ArrayList<String>();
        Matcher m = Pattern.compile("mul\\(\\d+,\\d+\\)").matcher(text);

        while (m.find()) {
            muls.add(m.group());
        }

        return muls;
    }

    public int[] getMulValues(String mul) {
        Matcher m = Pattern.compile("\\d+").matcher(mul);
        int[] values = new int[2];

        for (int i = 0; i < 2; i++) {
            m.find();
            values[i] = Integer.parseInt(m.group());
        }

        return values;
    }

    public int getMulResult(String mul) {
        int[] values = getMulValues(mul);
        return values[0] * values[1];
    }

    public int getMulsResult(String text) {
        ArrayList<String> muls = getMulsFromText(text);
        int result = 0;

        for (String mul : muls) {
            result += getMulResult(mul);
        }

        return result;
    }
}