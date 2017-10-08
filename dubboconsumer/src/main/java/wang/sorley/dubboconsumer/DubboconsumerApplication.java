package wang.sorley.dubboconsumer;
//import wang.sorley.demo.service.DubboDemoService;
//import org.springframework.context.support.ClassPathXmlApplicationContext;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class DubboconsumerApplication {

	public static void main(String[] args) {
//    	ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext(
//    			new String[]{"spring/application.xml"});
//    	context.start();
//		DubboDemoService service = (DubboDemoService)context.getBean("demoService");
//		String result = service.sayHello("sorley");
//		System.out.println(result);
//		try {
//			System.in.read();
//		}catch (Exception e){
//			e.printStackTrace();
//		}
		SpringApplication.run(DubboconsumerApplication.class, args);
	}
}
