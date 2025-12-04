package eu.houli.climbTheLeaderboard;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

public class ClimbTheLeaderboard
{
    public static void main( String[] args )
    {
        try {
            BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
            BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

            // Get the first line with the rankedCount
            Integer.parseInt(bufferedReader.readLine().trim());

            List<Integer> ranked = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                .map(Integer::parseInt)
                .collect(toList());

            // Get the playerScoreCount number from the next line
            Integer.parseInt(bufferedReader.readLine().trim());

            List<Integer> player = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                .map(Integer::parseInt)
                .collect(toList());

            List<Integer> result = ClimbTheLeaderboard.climbingLeaderboard(ranked, player);

            bufferedWriter.write(
                result.stream()
                    .map(Object::toString)
                    .collect(joining("\n"))
                + "\n"
            );

            bufferedReader.close();
            bufferedWriter.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /*
     * Complete the 'climbingLeaderboard' function below.
     *
     * The function is expected to return an INTEGER_ARRAY.
     * The function accepts following parameters:
     *  1. INTEGER_ARRAY ranked
     *  2. INTEGER_ARRAY player
     */
    public static List<Integer> climbingLeaderboard(List<Integer> ranked, List<Integer> player) {
        List<Integer> rankedDeDuped = ranked.stream().distinct().collect(toList());
        List<Integer> result = new ArrayList<>(player.size()); // we know the size already, so we can pre-size it
        int highIndex = rankedDeDuped.size() - 1;
        for (int p : player) {
            // [100, 50, 40, 20, 10]
            if (isFirst(p, rankedDeDuped)) {
                result.add(1);
                continue;
            }

            if(isLast(p, rankedDeDuped)) {
                result.add(rankedDeDuped.size()+1);
                continue;
            }

            highIndex = binarySearch(p, rankedDeDuped, highIndex);
            result.add(highIndex + 1);
        }
        return result;
    }

    public static List<Integer> doublePointer(List<Integer> ranked, List<Integer> player) {
        List<Integer> rankedDeDuped = ranked.stream().distinct().collect(toList());

        int playerIdx = 0;
        int rankedIdx = rankedDeDuped.size()-1;

        List<Integer> result = new ArrayList<>();
        while (playerIdx < player.size()) {
            if (rankedIdx == 0) {
                if (player.get(playerIdx) < rankedDeDuped.get(0)) {
                    result.add(2);
                } else {
                    result.add(1);
                }
                playerIdx++;
                continue;
            }

            int playerScore = player.get(playerIdx);
            int rankScore = rankedDeDuped.get(rankedIdx);

            if (playerScore < rankScore) {
                // add 1 for offset and 1 for being behind the current ranking
                result.add(rankedIdx + 2);
                playerIdx++;
                continue;
            }

            if (playerScore == rankScore) {
                // add 1 for offset only
                result.add(rankedIdx + 1);
                playerIdx++;
                continue;
            }

            if (playerScore > rankScore) {
                rankedIdx--;
                continue;
            }
        }
        return result;
    }

    private static boolean isFirst(int playerScore, List<Integer> ranked) {
        return playerScore >= ranked.get(0);
    }

    private static boolean isLast(int playerScore, List<Integer> ranked) {
        return playerScore < ranked.get(ranked.size()-1);
    }

    public static int binarySearch(int searchFor, List<Integer> collection, int high) {
        int low = 0;
        while (low <= high) {
            int mid = (low + high) / 2;
            int value = collection.get(mid);

            if (searchFor == value) {
                return mid;
            }

            if (searchFor < value) {
                // worse score - move right
                low = mid + 1;
            } else {
                // better score - move left
                high = mid - 1;
            }
        }
        return low;
    }
}
