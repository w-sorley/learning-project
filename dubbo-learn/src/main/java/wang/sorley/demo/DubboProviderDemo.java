package wang.sorley.demo;

import org.springframework.context.support.ClassPathXmlApplicationContext;


public class DubboProviderDemo {
    public static void main(String[] args) {
        ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext(
                new String[]{"spring/application.xml"});
        context.start();
        try {
            System.in.read();
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
