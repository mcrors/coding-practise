package eu.houli.magicSquare;

import java.io.*;
import java.util.*;
import java.util.stream.*;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'formingMagicSquare' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts 2D_INTEGER_ARRAY s as parameter.
     */

    private static final List<List<Integer>> preComputedMagicSquares = List.of(
        List.of(2, 7, 6, 1, 8, 3, 4, 9, 5),
        List.of(4, 9, 2, 7, 6, 1, 8, 3, 5),
        List.of(8, 3, 4, 9, 2, 7, 6, 1, 5),
        List.of(6, 1, 8, 3, 4, 9, 2, 7, 5),
        List.of(8, 1, 6, 7, 2, 9, 4, 3, 5),
        List.of(4, 3, 8, 1, 6, 7, 2, 9, 5),
        List.of(2, 9, 4, 3, 8, 1, 6, 7, 5),
        List.of(6, 7, 2, 9, 4, 3, 8, 1, 5)
    );

    private static final List<List<Integer>> unspoolOrder = List.of(
        List.of(0, 0), List.of(0, 1), List.of(0, 2),
        List.of(1, 2), List.of(2, 2), List.of(2, 1),
        List.of(2, 0), List.of(1, 0), List.of(1, 1)
    );

    private static List<Integer> unspoolInputSquare(List<List<Integer>> input) {
        List<Integer> result = new ArrayList<>();
        for(List<Integer> indexes : unspoolOrder) {
            int i = indexes.get(0);
            int j = indexes.get(1);
            result.add(input.get(i).get(j));
        }
        return result;
    }

    private static int difference(List<Integer> magicSquare, List<Integer> input) {
        int result = 0;
        for (int i = 0; i < 9; i++) {
            result += Math.abs(magicSquare.get(i) - input.get(i));
        }
        return result;
    }

    public static int formingMagicSquare(List<List<Integer>> s) {
        List<Integer> input = unspoolInputSquare(s);
        int result = 100;

        for (List<Integer> magicSquare : preComputedMagicSquares) {
            int latest = difference(magicSquare, input);
            result = latest < result ? latest : result;
        }

        return result;
    }
}

public class MagicSquare {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        List<List<Integer>> s = new ArrayList<>();

        IntStream.range(0, 3).forEach(i -> {
            try {
                s.add(
                    Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                        .map(Integer::parseInt)
                        .collect(toList())
                );
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        int result = Result.formingMagicSquare(s);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
