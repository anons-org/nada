

import java.util.ArrayList;
import java.util.List;

public class TestDemo1 {

	public static void main(String[] a) {
		List<Integer> x = new ArrayList<>();
		x.add(1);
		Test t = new Test();
		t.add(1, 2);
		// System.out.printf("%d", add(1,2));
	}

	public int add(int a, int b) {
		return a + b;
	}
}
