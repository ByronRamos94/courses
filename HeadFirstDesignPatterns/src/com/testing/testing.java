package com.testing;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class testing {
    public static List<List<String>> performWarehouseQueries(List<List<String>> query) {
        // Write your code here
        Queue<String> queue = new LinkedList<>();
        List<List<String>> response = new ArrayList<List<String>>();
        for (List<String> q : query) {
            if(q.get(0).equals("INSERT")){
                queue.add(q.get(1));
            } else if (q.get(0).equals("SHIP")){
                List<String> respItem = new ArrayList<>();
                if(queue.size() > 3){
                    for (int i = 0; i < 3; i++) {
                        respItem.add(queue.remove());
                    }
                } else {
                    respItem.add("NA");
                }
            }
        }
        return response;
    }
}
