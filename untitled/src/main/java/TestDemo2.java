

import java.util.ArrayList;
import java.util.List;

public class TestDemo2 {

	public static void main(String[] a) {
		List<Integer> x = new ArrayList<>();
		x.add(1);
		TestDemo2 t = new TestDemo2();
		t.add(1, 2);
		// System.out.printf("%d", add(1,2));
	}

	public int add(int a, int b) {
		return a + b;
	}
}
