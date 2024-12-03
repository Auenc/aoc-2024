import java.util.ArrayList;

public class Tests {
    public static void main(String[] args) {
        testGetMuls();
        testGetMulValues();
        testGetMulsResult();
    }

    public static void testGetMuls() {
        Day3 day3 = new Day3();
        ArrayList<String> muls = day3
                .getMulsFromText("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))");

        if (muls.size() != 4) {
            throw new AssertionError("Expected 4 muls, got " + muls.size());
        }
        if (!muls.get(0).equals("mul(2,4)")) {
            throw new AssertionError("Expected mul(2,4), got " + muls.get(0));
        }
        if (!muls.get(1).equals("mul(5,5)")) {
            throw new AssertionError("Expected mul(5,5), got " + muls.get(1));
        }
        if (!muls.get(2).equals("mul(11,8)")) {
            throw new AssertionError("Expected mul(11,8), got " + muls.get(2));
        }
        if (!muls.get(3).equals("mul(8,5)")) {
            throw new AssertionError("Expected mul(8,5), got " + muls.get(3));
        }
    }

    public static void testGetMulValues() {
        Day3 day3 = new Day3();
        {
            int[] values = day3.getMulValues("mul(2,4)");

            if (values[0] != 2) {
                throw new AssertionError("Expected 2, got " + values[0]);
            }
            if (values[1] != 4) {
                throw new AssertionError("Expected 4, got " + values[1]);
            }
        }
        {
            int[] values = day3.getMulValues("mul(100,5000)");

            if (values[0] != 100) {
                throw new AssertionError("Expected 100, got " + values[0]);
            }
            if (values[1] != 5000) {
                throw new AssertionError("Expected 5000, got " + values[1]);
            }
        }
    }

    public static void testGetMulsResult() {
        Day3 day3 = new Day3();
        int result = day3.getMulsResult("mul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))");

        if (result != 161) {
            throw new AssertionError("Expected 2 * 4 + 5 * 5 + 11 * 8 + 8 * 5, got " + result);
        }
    }
}
