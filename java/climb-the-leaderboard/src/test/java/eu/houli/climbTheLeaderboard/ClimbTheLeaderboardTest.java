package eu.houli.climbTheLeaderboard;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;


public class ClimbTheLeaderboardTest
{

   @ParameterizedTest
   @MethodSource("rankingTestCases")
   public void testClimbTheLeaderboard(
        List<Integer> ranked, List<Integer> playerScore, List<Integer> expected
    ) {
       List<Integer> result = ClimbTheLeaderboard.climbingLeaderboard(ranked, playerScore);

       assertEquals(expected, result);
    }

    @ParameterizedTest
    @MethodSource("rankingTestCases")
    public void testDoublePointed(
        List<Integer> ranked, List<Integer> playerScore, List<Integer> expected
    ) {
        List<Integer> result = ClimbTheLeaderboard.doublePointer(ranked, playerScore);
        assertEquals(expected, result);
    }


    private static Stream<Arguments> rankingTestCases() {
        return Stream.of(
                    Arguments.of(
                        List.of(100, 90, 90, 80),  // ranked
                        List.of(70, 80, 105),  // playerScore
                        List.of(4, 3, 1)   // expected
                    ),
                    Arguments.of(
                        List.of(100, 100, 50, 40, 40, 20, 10),
                        List.of(5, 25, 50, 120),
                        List.of(6, 4, 2, 1)
                    ),
                    Arguments.of(
                        List.of(100, 90, 80, 70),
                        List.of(60, 75, 95, 105),
                        List.of(5, 4, 2, 1)
                    ),
                    Arguments.of(
                        List.of(100, 90, 80),
                        List.of(85, 86, 87),
                        List.of(3, 3, 3)
                    )
            );
    }

    @ParameterizedTest
    @MethodSource("binarySearchTestCases")
    public void testBinarySearch(List<Integer> collection, int searchFor, int expected) {
        int result = ClimbTheLeaderboard.binarySearch(searchFor, collection, collection.size()-1);

        assertEquals(expected, result);

    }

    private static Stream<Arguments> binarySearchTestCases() {
        return Stream.of(
                Arguments.of(
                    List.of(100, 90, 80), // collection
                    80, // searchFor
                    2 // expected
                ),
                Arguments.of(
                    List.of(100, 50, 40, 20, 10),
                    25,
                    3
                ),
                Arguments.of(
                    List.of(100, 95, 70, 65, 55, 50, 45, 35, 20, 10, 5),
                    22,
                    8
                ),
                Arguments.of(
                    List.of(100, 50, 40, 20, 10),
                    75,
                    1
                )
            );
    }
}
