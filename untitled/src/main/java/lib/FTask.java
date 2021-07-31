package lib;

public class FTask {

	private IFTask task;

	public FTask(IFTask e) {
		// TODO Auto-generated constructor stub
		task = e;
	}

	public void start() {
		task.run(12);
	}

}
