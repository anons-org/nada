

public class TestDemo4 {

//	private String name;
//	public String userName;

	public static void main(String[] a) {
		int l = 10, r = 2;

		// IEEE754

		int farb = 99999999;
		// 常量池中没有
		// float far = 9.99999f;
		boolean farc = true;
		// long fard = 999999999999999999L;
		short fare = 111;
		// double fara = 99.22222222222222222222;

		String s = "世界您好！";
		// System.out.println(s);

		fare = 22;
		farc = false;

		r = add(l, r);
		Test test = new Test();
		test.add(l, r);
		test.add();
		// test.name="name";
		// test.userName = "uname";

//		
//		
//		new Thread(()->{
//			int xxxx = 88;
//			System.out.println("测试线程");
//			
//		}).start();

//		
//    	ITestDynMethod attr = (x)->{
//			System.out.println(x[0]);
//			return "";
//    	};
//
//    	attr.act("测试下",11);

//		ITestDynMethod lmd = new ITestDynMethod() {
//			
//			@Override
//			public Object act(Object... arg) {
//				// TODO Auto-generated method stub
//				System.out.println(arg[0]);
//				return null;
//			}
//		};
//	   lmd.act("测试");

//		Runnable race1 = new Runnable() {  
//		    @Override  
//		    public void run() {  
//		        System.out.println("Hello world !");  
//		    }  
//		};  

		// System.out.printf("%d", add(1,2));
	}

	public static int add(int a, int b) {
		return a + b;
	}

	public int add() {
		return 11;
	}

}
