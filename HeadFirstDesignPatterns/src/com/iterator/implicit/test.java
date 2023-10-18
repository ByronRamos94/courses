package com.iterator.implicit;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class test {

    public static void main(String[] args) {
        String s = "|**|*|*";
        List<Integer> startIndices = new ArrayList<>();
        List<Integer> endIndices = new ArrayList<>();
        startIndices.add(1);
        startIndices.add(1);
        endIndices.add(5);
        endIndices.add(6);

        // Write your code here
        String firstSubstring = s.substring(startIndices.get(0), endIndices.get(0));
        String secondSubString = s.substring(startIndices.get(1), endIndices.get(1));
        char[] charArray = firstSubstring.toCharArray();
        System.out.println(firstSubstring);
        System.out.println(secondSubString);
        Integer counter = 0;
        boolean increasing = false;
        char delimiter = '|';
        char item = '*';
        ArrayList<Integer> response = new ArrayList<>();

        for (int i = 0; i < firstSubstring.length(); i++) {
            if (firstSubstring.charAt(i) == delimiter) {
                if (!increasing) {
                    increasing = true;
                } else {
                    if (counter != 0) {
                        response.add(counter);
                    }
                }
            }
            if (firstSubstring.charAt(i) == item) {
                if (increasing) {
                    counter++;
                }
            }
        }
        counter = 0;
        for (int j = 0; j < secondSubString.length(); j++) {
            if (secondSubString.charAt(j) == delimiter) {
                if (!increasing) {
                    increasing = true;
                } else {
                    if (counter != 0) {
                        response.add(counter);
                    }
                }
            }
            if (secondSubString.charAt(j) == item) {
                if (increasing) {
                    counter++;
                }
            }
        }
        System.out.println(Arrays.toString(response.toArray()));

    }

}
