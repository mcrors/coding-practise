package eu.houli.magicSquare;

import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.summingInt;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'formingMagicSquare' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts 2D_INTEGER_ARRAY s as parameter.
     */

    private final List<List<List<Integer>>> preComputedMagicSquares = List.of(
        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        ),

        List.of(
            List.of(), // row 0
            List.of(), // row 1
            List.of() // row 1
        )
    );

    public static int formingMagicSquare(List<List<Integer>> s) {
    // Write your code here
        return 0;
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
