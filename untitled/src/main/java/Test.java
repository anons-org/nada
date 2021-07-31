
import lib.FTask;

public class Test {

	static int age = 100;

	static {
		age=200;
	}
	
	public static void enen(int i, int c, String s) {

	}

	public static void main(String[]  args){
		String s="打印输出测试";
		System.out.print(s);
	}

	native public  void testOk ();
	public   static void main11(String[] args){
	

		
		String testString = "a中文测试yo";
		Double abcc = 128.44444444444444444444444444444444444444444;
	//	BigDecimal abcca = new BigDecimal(128.4444444444444444444444444444444444444444);
		int inttest = 989899898;
		float fltest = 1.555555555566666666666666666666666666666666666f;
		Double abc = 1.22233333333333333333333333;
		long xxxxx = 888888888888888888L;
		boolean booltest = true;
		
		
		
		
		Object ab2c = 125;//"QQ184377367";
		
		
		
		
		long start = System.currentTimeMillis();
		
		//System.exit(1000);
		
		for(int x2=0; x2 < 1000 ;x2++) {
			new FTask((x) -> {
				for(int x1=0; x1 < 1000 ;x1++) {
					//System.out.println(x1);
				
				}
				System.gc();
			});
			
		}
		//System.nanoTime();
//		
//		long end = System.currentTimeMillis();
//		start = end - start;
//		System.out.println("Coroutine execution completion time <ms>");
//		System.out.println(start );
		
		//System.out.println( start / 1000 );
		
		
		

//    	IFTask attr = (x)->{
//			System.out.println(x);
//    	};
		
		
//		long start = System.currentTimeMillis();
//		System.out.println("测试");
//		for (int x = 0; x < l; x++) {
//			System.out.println(x);
//		}
//		long end = System.currentTimeMillis();
//		start = end - start;
//		System.out.println(start);
//		
//		System.out.println( start / 1000 );
//		
//		for(int x=88; x > 0 ;x--) {
//			System.out.print(11);
//		}

//		
//		//IEEE754
//	
//		int 		farb 	= 99999999;
//		//常量池中没有
//		//float 		far		= 9.99999f;
//		boolean    	farc	= true;
//		//long		fard	= 999999999999999999L;
//		short		fare	= 111;
//		//double 		fara 	= 99.22222222222222222222;
//		
//		
//		
//		String s ="世界您好！";
//	//	System.out.println(s);
//		
//		
//		fare=22;
//		farc = false;
//	
//		r = add(l, r);
//		Test test = new Test();
//		test.add(l, r);
//		test.add();
//		//test.name="name";
//		//test.userName = "uname";
//	

//		
//		
//		new Thread(()->{
//			int xxxx = 88;
//			System.out.println("测试线程");
//			
//		}).start();
//		

//		new Thread(()->{
//			int xxxx = 88;
//			System.out.println("测试线程");
//			
//		}).start();
//		

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
