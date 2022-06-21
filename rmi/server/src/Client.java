
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;

public class Client {

    private Client() {}

    public static void main(String[] args) {

		String host = (args.length < 1) ? null : args[0];
		try {
		    Registry registry = LocateRegistry.getRegistry(host);
		    Hello stub = (Hello) registry.lookup("Hello");
		    String response = stub.sayHello();
		    System.out.println("response: " + response);
		    MyTask myTask = new MyTask();
		    System.out.println("client status1 = " + myTask.getStatus());
		    Object o = stub.execute(myTask);
		    System.out.println("o: " + o);
		    System.out.println("client status2 = " + myTask.getStatus());
		    
		} catch (Exception e) {
		    System.err.println("Client exception: " + e.toString());
		    e.printStackTrace();
		}
    }
}
