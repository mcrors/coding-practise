package eu.houli.magicSquare;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;

/**
 * Unit test for simple App.
 */
public class MagicSquareTest
{

   @ParameterizedTest
   @MethodSource("rankingTestCases")
   public void testMagicSquare(List<List<Integer>> magicSquare, int expected) {
        int result = Result.formingMagicSquare(magicSquare);
        assertEquals(expected, result);
    }


    private static Stream<Arguments> rankingTestCases() {
        return Stream.of(
                    Arguments.of(
                        List.of(
                            List.of(5, 3, 4),
                            List.of(1, 5, 8),
                            List.of(6, 4, 2)
                        ),
                        7
                    ),
                    Arguments.of(
                        List.of(
                            List.of(4, 9, 2),
                            List.of(3, 5, 7),
                            List.of(8, 1, 5)
                        ),
                        1
                    ),
                    Arguments.of(
                        List.of(
                            List.of(4, 5, 8),
                            List.of(2, 4, 1),
                            List.of(1, 9, 7)
                        ),
                        14
                    )
            );
    }
}
