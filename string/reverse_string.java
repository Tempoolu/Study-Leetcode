package string;

import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedList;
import java.util.List;

import try_java.java_basic.list;

public class reverse_string {
    public static void main(String[] args) {

        reverse_string.reverse2("abcdefghijklmn", 3);

        List<Character> tt = Arrays.asList('a', 'b', 'c', 'd');
        
        tt = rev(tt);
    }


    public static List<Character> rev(List<Character> li) {
        int len = li.size();
        for (int i=0; i<len/2; i++) {
            char temp = li.get(i);
            li.set(i, li.get(len-1-i));
            li.set(len-1-i, temp);
        }
        return li;
    }

    public static String reverse2(String s, int k) {
        int len = s.length();
        char[] ch = s.toCharArray();
        List<Character> li = new LinkedList<>(), pre = new LinkedList<>(), res = new LinkedList<>();;

        for (Character c : ch) {
            li.add(c);
        }

        pre = li.subList(0, (len/2/k)*2*k);
        
        
        for (int i=0; i<pre.size(); i+=2*k) {
            List<Character> subLi = new LinkedList<>();
            subLi = reverse_string.rev(pre.subList(i, i+k));
            res.addAll(subLi);
            for (int j=i+k; j<i+2*k; j++) {
                res.add(pre.get(j));
            }
            
        }
        System.out.println(res);
        
        
        if (len/(2*k) < k) {

        } else if (len/(2*k) > k) {
            
        }


        return "aa";
    }


    
}
